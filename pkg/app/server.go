package app

import (
	"github.com/appscode/log"
	app "github.com/appscode/wheel/pkg/apis/app/v1beta1"
	"github.com/appscode/wheel/pkg/apiserver/endpoints"
	"golang.org/x/net/context"
	"k8s.io/helm/pkg/proto/hapi/chart"
	rls "k8s.io/helm/pkg/proto/hapi/services"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(app.RegisterReleaseServiceServer, &AppsServer{})
	endpoints.ProxyServerEndpoints.Register(app.RegisterReleaseServiceHandlerFromEndpoint)
}

const (
	WHEEL_ARCHIVE = "/tmp/wheel-archive/"
	DEFAULT_NS    = "default"
)

type AppsServer struct{}

var _ app.ReleaseServiceServer = &AppsServer{}

func (*AppsServer) ListReleases(req *app.ListReleasesRequest, srv app.ReleaseService_ListReleasesServer) error {
	rlc, err := getReleaseServiceClient()
	if err != nil {
		return err
	}

	if req.Namespace == "" {
		req.Namespace = DEFAULT_NS
	}

	listReq := rls.ListReleasesRequest{
		Filter:      req.Filter,
		Limit:       req.Limit,
		Namespace:   req.Namespace,
		Offset:      req.Offset,
		SortBy:      rls.ListSort_SortBy(rls.ListSort_SortBy_value[req.SortBy.String()]),
		SortOrder:   rls.ListSort_SortOrder(rls.ListSort_SortOrder_value[req.SortOrder.String()]),
		StatusCodes: req.StatusCodes,
	}

	listClient, err := rlc.ListReleases(newContext(), &listReq)
	if err != nil {
		return err
	}

	listRes, err := listClient.Recv()
	if err != nil {
		return err
	}

	res := app.ListReleasesResponse{
		Count:    listRes.Count,
		Next:     listRes.Next,
		Releases: listRes.Releases,
		Total:    listRes.Total,
	}

	srv.Send(&res)
	return nil
}

// GetReleasesStatus retrieves status information for the specified release.
func (*AppsServer) GetReleaseStatus(ctx context.Context, req *app.GetReleaseStatusRequest) (*app.GetReleaseStatusResponse, error) {
	rlc, err := getReleaseServiceClient()
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

	return &app.GetReleaseStatusResponse{
		Name:      statusRes.Name,
		Namespace: statusRes.Namespace,
		Info:      statusRes.Info,
	}, nil
}

// GetReleaseContent retrieves the release content (chart + value) for the specified release.
func (*AppsServer) GetReleaseContent(ctx context.Context, req *app.GetReleaseContentRequest) (*app.GetReleaseContentResponse, error) {
	rlc, err := getReleaseServiceClient()
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

	return &app.GetReleaseContentResponse{
		Release: contentRes.Release,
	}, nil
}

// UpdateRelease updates release content.
func (*AppsServer) UpdateRelease(ctx context.Context, req *app.UpdateReleaseRequest) (*app.UpdateReleaseResponse, error) {
	rlc, err := getReleaseServiceClient()
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

	log.Infoln(req)

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

	return &app.UpdateReleaseResponse{
		Release: updateRes.Release,
	}, nil
}

// InstallRelease requests installation of a chart as a new release.
func (*AppsServer) InstallRelease(ctx context.Context, req *app.InstallReleaseRequest) (*app.InstallReleaseResponse, error) {
	rlc, err := getReleaseServiceClient()
	if err != nil {
		return nil, err
	}

	if req.Namespace == "" {
		req.Namespace = DEFAULT_NS
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

	log.Infoln(req)

	installReq := rls.InstallReleaseRequest{
		Name:         req.Name,
		Timeout:      req.Timeout,
		Chart:        req.Chart,
		DisableHooks: req.DisableHooks,
		DryRun:       req.DryRun,
		Values:       req.Values,
		Wait:         req.Wait,
		Namespace:    req.Namespace,
		ReuseName:    req.ReuseName,
	}

	installRes, err := rlc.InstallRelease(newContext(), &installReq)
	if err != nil {
		return nil, err
	}

	return &app.InstallReleaseResponse{
		Release: installRes.Release,
	}, nil
}

// UninstallRelease requests deletion of a named release.
func (*AppsServer) UninstallRelease(ctx context.Context, req *app.UninstallReleaseRequest) (*app.UninstallReleaseResponse, error) {
	rlc, err := getReleaseServiceClient()
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

	return &app.UninstallReleaseResponse{
		Release: uninstallRes.Release,
		Info:    uninstallRes.Info,
	}, nil
}

// GetVersion returns the current version of the server.
func (*AppsServer) GetVersion(ctx context.Context, req *app.GetVersionRequest) (*app.GetVersionResponse, error) {
	rlc, err := getReleaseServiceClient()
	if err != nil {
		return nil, err
	}

	versionReq := rls.GetVersionRequest{}

	versionRes, err := rlc.GetVersion(newContext(), &versionReq)
	if err != nil {
		return nil, err
	}

	return &app.GetVersionResponse{
		Version: versionRes.Version,
	}, nil
}

// RollbackRelease rolls back a release to a previous version.
func (*AppsServer) RollbackRelease(ctx context.Context, req *app.RollbackReleaseRequest) (*app.RollbackReleaseResponse, error) {
	rlc, err := getReleaseServiceClient()
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

	return &app.RollbackReleaseResponse{
		Release: rollbackRes.Release,
	}, nil
}

// ReleaseHistory retrieves a releasse's history.
func (*AppsServer) GetHistory(ctx context.Context, req *app.GetHistoryRequest) (*app.GetHistoryResponse, error) {
	rlc, err := getReleaseServiceClient()
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

	return &app.GetHistoryResponse{
		Releases: historyRes.Releases,
	}, nil
}

// RunReleaseTest executes the tests defined of a named release
func (*AppsServer) RunReleaseTest(req *app.TestReleaseRequest, srv app.ReleaseService_RunReleaseTestServer) error {
	rlc, err := getReleaseServiceClient()
	if err != nil {
		return err
	}

	testReq := rls.TestReleaseRequest{
		Name:    req.Name,
		Cleanup: req.Cleanup,
		Timeout: req.Timeout,
	}

	testClient, err := rlc.RunReleaseTest(newContext(), &testReq)
	if err != nil {
		return err
	}

	testRes, err := testClient.Recv()
	if err != nil {
		return err
	}

	res := app.TestReleaseResponse{
		Msg:    testRes.Msg,
		Status: testRes.Status,
	}

	srv.Send(&res)
	return nil
}
