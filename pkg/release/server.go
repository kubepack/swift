package release

import (
	stringz "github.com/appscode/go/strings"
	proto "github.com/appscode/swift/pkg/apis/swift/v2"
	"github.com/appscode/swift/pkg/connectors"
	"github.com/appscode/swift/pkg/extpoints"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/helm/pkg/proto/hapi/chart"
	hrls "k8s.io/helm/pkg/proto/hapi/release"
	rls "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/version"
)

type Server struct {
	ClientFactory extpoints.Connector
}

var _ proto.ReleaseServiceServer = &Server{}

// NewContext creates a versioned context.
func newContext() context.Context {
	md := metadata.Pairs("x-helm-api-client", version.GetVersion())
	return metadata.NewOutgoingContext(context.TODO(), md)
}

func (s *Server) SummarizeReleases(ctx context.Context, req *proto.SummarizeReleasesRequest) (*proto.SummarizeReleasesResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)
	listReq := rls.ListReleasesRequest{
		Filter:      req.Filter,
		Limit:       req.Limit,
		Namespace:   req.Namespace,
		Offset:      req.Offset,
		SortBy:      rls.ListSort_SortBy(rls.ListSort_SortBy_value[req.SortBy.String()]),
		SortOrder:   rls.ListSort_SortOrder(rls.ListSort_SortOrder_value[req.SortOrder.String()]),
		StatusCodes: []hrls.Status_Code{},
	}

	if req.All { // list all releases
		listReq.StatusCodes = []hrls.Status_Code{
			hrls.Status_UNKNOWN,
			hrls.Status_DEPLOYED,
			hrls.Status_DELETED,
			hrls.Status_SUPERSEDED,
			hrls.Status_FAILED,
			hrls.Status_DELETING,
			hrls.Status_PENDING_INSTALL,
			hrls.Status_PENDING_UPGRADE,
			hrls.Status_PENDING_ROLLBACK,
		}
	} else { // convert status(string) to status-code(int32)
		for _, status := range req.StatusCodes {
			if val, ok := hrls.Status_Code_value[status]; ok {
				listReq.StatusCodes = append(listReq.StatusCodes, hrls.Status_Code(val))
			}
		}
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	listClient, err := rlc.ListReleases(newContext(), &listReq)
	if err != nil {
		return nil, err
	}

	listRes, err := listClient.Recv()
	if err != nil {
		return nil, err
	}

	var releases []*proto.ReleaseSummary

	for _, item := range listRes.Releases {
		releases = append(releases, &proto.ReleaseSummary{
			Namespace:     item.Namespace,
			Name:          item.Name,
			Info:          item.Info,
			Version:       item.Version,
			Config:        item.Config,
			ChartMetadata: item.Chart.Metadata,
		})
	}

	return &proto.SummarizeReleasesResponse{
		Releases: releases,
	}, nil

}

// GetReleasesStatus retrieves status information for the specified release.
func (s *Server) GetReleaseStatus(ctx context.Context, req *proto.GetReleaseStatusRequest) (*proto.GetReleaseStatusResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	statusReq := rls.GetReleaseStatusRequest{
		Name:    req.Name,
		Version: req.Version,
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	statusRes, err := rlc.GetReleaseStatus(newContext(), &statusReq)
	if err != nil {
		return nil, err
	}

	return &proto.GetReleaseStatusResponse{
		Name:      statusRes.Name,
		Namespace: statusRes.Namespace,
		Info:      statusRes.Info,
	}, nil
}

// GetReleaseContent retrieves the release content (chart + value) for the specified release.
func (s *Server) GetReleaseContent(ctx context.Context, req *proto.GetReleaseContentRequest) (*proto.GetReleaseContentResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	contentReq := rls.GetReleaseContentRequest{
		Name:    req.Name,
		Version: req.Version,
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	contentRes, err := rlc.GetReleaseContent(newContext(), &contentReq)
	if err != nil {
		return nil, err
	}

	// Format release config and values to JSON string
	// If error found, skip formatting
	if req.FormatValuesAsJson {
		if config, err := yaml.ToJSON([]byte(contentRes.Release.Config.Raw)); err == nil {
			contentRes.Release.Config.Raw = string(config)
		}
		if config, err := yaml.ToJSON([]byte(contentRes.Release.Chart.Values.Raw)); err == nil {
			contentRes.Release.Chart.Values.Raw = string(config)
		}
	}

	return &proto.GetReleaseContentResponse{
		Release: contentRes.Release,
	}, nil
}

// UpdateRelease updates release content.
func (s *Server) UpdateRelease(ctx context.Context, req *proto.UpdateReleaseRequest) (*proto.UpdateReleaseResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	if req.Values == nil { // (req.Values == nil) causes render error
		req.Values = &chart.Config{}
	}

	if req.Chart == nil {
		repo := chartInfo{
			ChartURL:           req.ChartUrl,
			CaBundle:           req.CaBundle,
			Username:           req.Username,
			Password:           req.Password,
			Token:              req.Token,
			ClientCertificate:  req.ClientCertificate,
			ClientKey:          req.ClientKey,
			InsecureSkipVerify: req.InsecureSkipVerify,
		}
		req.Chart, err = prepareChart(repo, req.Values)
		if err != nil {
			return nil, err
		}
	}

	updateReq := rls.UpdateReleaseRequest{
		Name:         req.Name,
		Timeout:      req.Timeout,
		Chart:        req.Chart,
		DisableHooks: req.DisableHooks,
		DryRun:       req.DryRun,
		Force:        req.Force,
		Recreate:     req.Recreate,
		ResetValues:  req.ResetValues,
		ReuseValues:  req.ReuseValues,
		Values:       req.Values,
		Wait:         req.Wait,
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	updateRes, err := rlc.UpdateRelease(newContext(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateReleaseResponse{
		Release: updateRes.Release,
	}, nil
}

// InstallRelease requests installation of a chart as a new release.
func (s *Server) InstallRelease(ctx context.Context, req *proto.InstallReleaseRequest) (*proto.InstallReleaseResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	if req.Values == nil { // (req.Values == nil) causes render error
		req.Values = &chart.Config{}
	}

	if req.Chart == nil {
		repo := chartInfo{
			ChartURL:           req.ChartUrl,
			CaBundle:           req.CaBundle,
			Username:           req.Username,
			Password:           req.Password,
			Token:              req.Token,
			ClientCertificate:  req.ClientCertificate,
			ClientKey:          req.ClientKey,
			InsecureSkipVerify: req.InsecureSkipVerify,
		}
		req.Chart, err = prepareChart(repo, req.Values)
		if err != nil {
			return nil, err
		}
	}

	installReq := rls.InstallReleaseRequest{
		Name:         req.Name,
		Timeout:      req.Timeout,
		Chart:        req.Chart,
		DisableHooks: req.DisableHooks,
		DryRun:       req.DryRun,
		Values:       req.Values,
		Wait:         req.Wait,
		Namespace:    stringz.Val(req.Namespace, core.NamespaceDefault),
		ReuseName:    req.ReuseName,
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	installRes, err := rlc.InstallRelease(newContext(), &installReq)
	if err != nil {
		return nil, err
	}

	return &proto.InstallReleaseResponse{
		Release: installRes.Release,
	}, nil
}

// UninstallRelease requests deletion of a named release.
func (s *Server) UninstallRelease(ctx context.Context, req *proto.UninstallReleaseRequest) (*proto.UninstallReleaseResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	uninstallReq := rls.UninstallReleaseRequest{
		Name:         req.Name,
		Timeout:      req.Timeout,
		DisableHooks: req.DisableHooks,
		Purge:        req.Purge,
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	uninstallRes, err := rlc.UninstallRelease(newContext(), &uninstallReq)
	if err != nil {
		return nil, err
	}

	return &proto.UninstallReleaseResponse{
		Release: uninstallRes.Release,
		Info:    uninstallRes.Info,
	}, nil
}

// GetVersion returns the current version of the server.
func (s *Server) GetVersion(ctx context.Context, req *proto.GetVersionRequest) (*proto.GetVersionResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	versionReq := rls.GetVersionRequest{}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	versionRes, err := rlc.GetVersion(newContext(), &versionReq)
	if err != nil {
		return nil, err
	}

	return &proto.GetVersionResponse{
		Version: versionRes.Version,
	}, nil
}

// RollbackRelease rolls back a release to a previous version.
func (s *Server) RollbackRelease(ctx context.Context, req *proto.RollbackReleaseRequest) (*proto.RollbackReleaseResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	rollbackReq := rls.RollbackReleaseRequest{
		Name:         req.Name,
		Timeout:      req.Timeout,
		DisableHooks: req.DisableHooks,
		DryRun:       req.DryRun,
		Wait:         req.Wait,
		Recreate:     req.Recreate,
		Force:        req.Force,
		Version:      req.Version,
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	rollbackRes, err := rlc.RollbackRelease(newContext(), &rollbackReq)
	if err != nil {
		return nil, err
	}

	return &proto.RollbackReleaseResponse{
		Release: rollbackRes.Release,
	}, nil
}

// ReleaseHistory retrieves a release's history.
func (s *Server) GetHistory(ctx context.Context, req *proto.GetHistoryRequest) (*proto.GetHistoryResponse, error) {
	ctx, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.ClientFactory.Close(ctx)

	historyReq := rls.GetHistoryRequest{
		Name: req.Name,
		Max:  req.Max,
	}

	rlc := rls.NewReleaseServiceClient(connectors.Connection(ctx))
	historyRes, err := rlc.GetHistory(newContext(), &historyReq)
	if err != nil {
		return nil, err
	}

	return &proto.GetHistoryResponse{
		Releases: historyRes.Releases,
	}, nil
}
