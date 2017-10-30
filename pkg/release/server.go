package release

import (
	stringz "github.com/appscode/go/strings"
	proto "github.com/appscode/swift/pkg/apis/swift/v2"
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
	return metadata.NewContext(context.TODO(), md)
}

func (s *Server) SummarizeReleases(ctx context.Context, req *proto.SummarizeReleasesRequest) (*proto.SummarizeReleasesResponse, error) {
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}
	listReq := rls.ListReleasesRequest{
		Filter:      req.Filter,
		Limit:       req.Limit,
		Namespace:   req.Namespace,
		Offset:      req.Offset,
		SortBy:      rls.ListSort_SortBy(rls.ListSort_SortBy_value[req.SortBy.String()]),
		SortOrder:   rls.ListSort_SortOrder(rls.ListSort_SortOrder_value[req.SortOrder.String()]),
		StatusCodes: []hrls.Status_Code{},
	}

	if len(req.StatusCodes) == 0 { // list all releases
		listReq.StatusCodes = []hrls.Status_Code{
			hrls.Status_UNKNOWN,
			hrls.Status_DEPLOYED,
			hrls.Status_DELETED,
			hrls.Status_SUPERSEDED,
			hrls.Status_FAILED,
			hrls.Status_DELETING,
		}
	} else { // convert status(string) to status-code(int32)
		for _, status := range req.StatusCodes {
			if val, ok := hrls.Status_Code_value[status]; ok {
				listReq.StatusCodes = append(listReq.StatusCodes, hrls.Status_Code(val))
			}
		}
	}

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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

	statusReq := rls.GetReleaseStatusRequest{
		Name:    req.Name,
		Version: req.Version,
	}

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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

	contentReq := rls.GetReleaseContentRequest{
		Name:    req.Name,
		Version: req.Version,
	}

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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

	if req.Values == nil { // (req.Values == nil) causes render error
		req.Values = &chart.Config{}
	}

	if req.Chart == nil {
		req.Chart, err = prepareChart(req.ChartUrl, req.Values)
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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

	if req.Values == nil { // (req.Values == nil) causes render error
		req.Values = &chart.Config{}
	}

	if req.Chart == nil {
		req.Chart, err = prepareChart(req.ChartUrl, req.Values)
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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

	uninstallReq := rls.UninstallReleaseRequest{
		Name:         req.Name,
		Timeout:      req.Timeout,
		DisableHooks: req.DisableHooks,
		Purge:        req.Purge,
	}

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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

	versionReq := rls.GetVersionRequest{}

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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

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
	rlc, err := s.ClientFactory.Connect(ctx)
	if err != nil {
		return nil, err
	}

	historyReq := rls.GetHistoryRequest{
		Name: req.Name,
		Max:  req.Max,
	}

	historyRes, err := rlc.GetHistory(newContext(), &historyReq)
	if err != nil {
		return nil, err
	}

	return &proto.GetHistoryResponse{
		Releases: historyRes.Releases,
	}, nil
}
