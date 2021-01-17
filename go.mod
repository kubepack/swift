module kubepack.dev/swift

go 1.15

require (
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/opentracing/opentracing-go v1.0.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
	gomodules.xyz/errors v0.0.0-20201104190405-077f059979fd
	gomodules.xyz/grpc-go-addons v0.2.1
	gomodules.xyz/runtime v0.0.0-20201104200926-d838b09dda8b
	gomodules.xyz/signals v0.0.0-20201104192641-f8f5c878d966
	gomodules.xyz/x v0.0.0-20201105065653-91c568df6331
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a
	google.golang.org/grpc v1.26.0
	k8s.io/api v0.18.9
	k8s.io/apimachinery v0.18.9
	k8s.io/client-go v0.18.9
	k8s.io/helm v2.14.1+incompatible
	kmodules.xyz/client-go v0.0.0-20210117055448-ac3622fe26e7
)

replace cloud.google.com/go => cloud.google.com/go v0.34.0

replace git.apache.org/thrift.git => github.com/apache/thrift v0.12.0

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.2+incompatible

replace github.com/grpc-ecosystem/go-grpc-middleware => github.com/gomodules/go-grpc-middleware v0.0.0-20180226223443-606e44dc6300

replace github.com/grpc-ecosystem/grpc-gateway => github.com/gomodules/grpc-gateway v1.3.1-ac

replace k8s.io/helm => github.com/appscode/helm v2.14.1-0.20190516122408-6f598715802a+incompatible

replace k8s.io/klog => k8s.io/klog v0.3.0

replace k8s.io/api => github.com/kmodules/api v0.18.10-0.20200922195318-d60fe725dea0

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200922195535-0c9a1b86beec

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.18.10-0.20200922195747-1bd1cc8f00d1

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.9

replace k8s.io/client-go => github.com/kmodules/k8s-client-go v0.18.10-0.20200922201634-73fedf3d677e

replace k8s.io/component-base => k8s.io/component-base v0.18.9

replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.19.0-alpha.0.0.20200922200158-8b13196d8dc4

replace k8s.io/utils => k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89

replace sigs.k8s.io/application => github.com/kmodules/application v0.8.4-0.20201117013009-57cb1e10e2ed
