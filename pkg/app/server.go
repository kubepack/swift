package app

import (
	"fmt"

	app "github.com/appscode/wheel/pkg/apis/app/v1beta1"
	"github.com/appscode/wheel/pkg/apiserver/endpoints"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(app.RegisterAppsServer, &AppsServer{})
	endpoints.ProxyServerEndpoints.Register(app.RegisterAppsHandlerFromEndpoint)

	//endpoints.GRPCServerEndpoints.Register(app.RegisterReleaseServiceServer, &AppsServer{})
	//endpoints.ProxyServerEndpoints.Register(app.RegisterReleaseServiceHandlerFromEndpoint)
}

type AppsServer struct{}

var _ app.AppsServer = &AppsServer{}
var _ app.ReleaseServiceServer = &AppsServer{}

func (*AppsServer) Hello(ctx context.Context, req *app.HelloRequest) (*app.HelloResponse, error) {
	return &app.HelloResponse{
		Greetings: fmt.Sprintf("Hello, %s.", req.Name),
	}, nil
}

func (*AppsServer) ListReleases(req *app.ListReleasesRequest, srv app.ReleaseService_ListReleasesServer) error {
	/*setupConnection()

	l := &listCmd{
		client: helm.NewClient(helm.Host(settings.TillerHost)),
	}

	res, err := l.run()
	if err != nil {
		return err
	}

	srv.Send(&res)*/
	return nil
}

// GetReleasesStatus retrieves status information for the specified release.
func (*AppsServer) GetReleaseStatus(ctx context.Context, req *app.GetReleaseStatusRequest) (*app.GetReleaseStatusResponse, error) {
	return nil, nil
}

// GetReleaseContent retrieves the release content (chart + value) for the specified release.
func (*AppsServer) GetReleaseContent(ctx context.Context, req *app.GetReleaseContentRequest) (*app.GetReleaseContentResponse, error) {
	return nil, nil

}

// UpdateRelease updates release content.
func (*AppsServer) UpdateRelease(ctx context.Context, req *app.UpdateReleaseRequest) (*app.UpdateReleaseResponse, error) {
	return nil, nil

}

// InstallRelease requests installation of a chart as a new release.
func (*AppsServer) InstallRelease(ctx context.Context, req *app.InstallReleaseRequest) (*app.InstallReleaseResponse, error) {
	return nil, nil

}

// UninstallRelease requests deletion of a named release.
func (*AppsServer) UninstallRelease(ctx context.Context, req *app.UninstallReleaseRequest) (*app.UninstallReleaseResponse, error) {
	return nil, nil

}

// GetVersion returns the current version of the server.
func (*AppsServer) GetVersion(ctx context.Context, req *app.GetVersionRequest) (*app.GetVersionResponse, error) {
	return nil, nil

}

// RollbackRelease rolls back a release to a previous version.
func (*AppsServer) RollbackRelease(ctx context.Context, req *app.RollbackReleaseRequest) (*app.RollbackReleaseResponse, error) {
	return nil, nil

}

// ReleaseHistory retrieves a releasse's history.
func (*AppsServer) GetHistory(ctx context.Context, req *app.GetHistoryRequest) (*app.GetHistoryResponse, error) {
	return nil, nil

}

// RunReleaseTest executes the tests defined of a named release
func (*AppsServer) RunReleaseTest(req *app.TestReleaseRequest, srv app.ReleaseService_RunReleaseTestServer) error {
	return nil

}
