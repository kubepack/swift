package app

import (
	"fmt"

	app "github.com/appscode/wheel/pkg/apis/app/v1beta1"
	"github.com/appscode/wheel/pkg/apiserver/endpoints"
	"golang.org/x/net/context"
	"k8s.io/helm/pkg/helm"
	"github.com/appscode/wheel/pkg/kube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	rls "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/version"


)

func init() {
	endpoints.GRPCServerEndpoints.Register(app.RegisterAppsServer, &AppsServer{})
	endpoints.ProxyServerEndpoints.Register(app.RegisterAppsHandlerFromEndpoint)

	endpoints.GRPCServerEndpoints.Register(app.RegisterReleaseServiceServer, &AppsServer{})
	endpoints.ProxyServerEndpoints.Register(app.RegisterReleaseServiceHandlerFromEndpoint)
}

type AppsServer struct{}

var _ app.AppsServer = &AppsServer{}
var _ app.ReleaseServiceServer = &AppsServer{}

func (*AppsServer) Hello(ctx context.Context, req *app.HelloRequest) (*app.HelloResponse, error) {
	return &app.HelloResponse{
		Greetings: fmt.Sprintf("Hello, %s.", req.Name),
	}, nil
}

// connect returns a grpc connection to tiller or error. The grpc dial options
// are constructed here.
func connect(ctx context.Context, host string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
	}
	opts = append(opts, grpc.WithInsecure())

	/*switch {
	case h.opts.useTLS:
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(h.opts.tlsConfig)))
	default:
		opts = append(opts, grpc.WithInsecure())
	}*/


	if conn, err = grpc.Dial(host, opts...); err != nil {
		return nil, err
	}
	return conn, nil
}

// NewContext creates a versioned context.
func NewContext() context.Context {
	md := metadata.Pairs("x-helm-api-client", version.GetVersion())
	return metadata.NewContext(context.TODO(), md)
}

func (*AppsServer) ListReleases(req *app.ListReleasesRequest, srv app.ReleaseService_ListReleasesServer) error {

	f := kube.NewKubeFactory()

	restClient, err := f.RESTClient()
	if err != nil {
		return err
	}

	config, err := f.ClientConfig()
	if err != nil {
		return err
	}

	operatorNamespace := "kube-system" //flag
	clientSet, err := f.ClientSet() //err
	operatorLabel := "app: helm"


	operatorPodList, err := clientSet.Core().Pods(operatorNamespace).List(
		metav1.ListOptions{
			LabelSelector: operatorLabel,
		},
	)
	if err != nil {
		return err
	}

	operatorPortNumber := 44134

	tunnel := newTunnel(restClient, config, operatorNamespace, operatorPodList.Items[0].Name, operatorPortNumber)
	if err := tunnel.forwardPort(); err != nil {
		return err
	}

	ctx := NewContext()
	host := fmt.Sprintf("127.0.0.1:%d",tunnel.Local)

	conn, err := connect(ctx, host)

	rlc := rls.NewReleaseServiceClient(conn)

	r := rls.ListReleasesRequest{
		Filter: req.Filter,
		Limit: req.Limit,
		Namespace: req.Namespace,
		Offset: req.Offset,
		SortBy:  rls.ListSort_SortBy(rls.ListSort_SortBy_value[req.SortBy.String()]) ,
		SortOrder: req.SortOrder,
		StatusCodes: req.SortOrder,
	}

	s, err := rlc.ListReleases(ctx, r)
	if err != nil {
		return nil, err
	}

	l := &listCmd{
		client: helm.NewClient(helm.Host(settings.TillerHost)),
	}

	res, err := l.run()
	if err != nil {
		return err
	}

	srv.Send(&res)
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
