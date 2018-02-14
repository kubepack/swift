#!/bin/bash
set -eou pipefail

echo "checking kubeconfig context"
kubectl config current-context || { echo "Set a context (kubectl use-context <context>) out of the following:"; echo; kubectl config get-contexts; exit 1; }
echo ""

# ref: https://stackoverflow.com/a/27776822/244009
case "$(uname -s)" in
    Darwin)
        curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.1.0/onessl-darwin-amd64
        chmod +x onessl
        export ONESSL=./onessl
        ;;

    Linux)
        curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.1.0/onessl-linux-amd64
        chmod +x onessl
        export ONESSL=./onessl
        ;;

    CYGWIN*|MINGW32*|MSYS*)
        curl -fsSL -o onessl.exe https://github.com/kubepack/onessl/releases/download/0.1.0/onessl-windows-amd64.exe
        chmod +x onessl.exe
        export ONESSL=./onessl.exe
        ;;
    *)
        echo 'other OS'
        ;;
esac

# http://redsymbol.net/articles/bash-exit-traps/
function cleanup {
    rm -rf $ONESSL ca.crt ca.key server.crt server.key
}
trap cleanup EXIT

# ref: https://stackoverflow.com/a/7069755/244009
# ref: https://jonalmeida.com/posts/2013/05/26/different-ways-to-implement-flags-in-bash/
# ref: http://tldp.org/LDP/abs/html/comparison-ops.html

export SWIFT_NAMESPACE=kube-system
export SWIFT_SERVICE_ACCOUNT=default
export SWIFT_ENABLE_RBAC=false
export SWIFT_RUN_ON_MASTER=0
export SWIFT_DOCKER_REGISTRY=appscode
export SWIFT_IMAGE_PULL_SECRET=
export SWIFT_UNINSTALL=0

show_help() {
    echo "swift.sh - install Ajax friendly Helm Tiller Proxy"
    echo " "
    echo "swift.sh [options]"
    echo " "
    echo "options:"
    echo "-h, --help                         show brief help"
    echo "-n, --namespace=NAMESPACE          specify namespace (default: kube-system)"
    echo "    --rbac                         create RBAC roles and bindings"
    echo "    --docker-registry              docker registry used to pull swift images (default: appscode)"
    echo "    --image-pull-secret            name of secret used to pull swift docker images"
    echo "    --run-on-master                run swift operator on master"
    echo "    --uninstall                    uninstall swift"
}

while test $# -gt 0; do
    case "$1" in
        -h|--help)
            show_help
            exit 0
            ;;
        -n)
            shift
            if test $# -gt 0; then
                export SWIFT_NAMESPACE=$1
            else
                echo "no namespace specified"
                exit 1
            fi
            shift
            ;;
        --namespace*)
            export SWIFT_NAMESPACE=`echo $1 | sed -e 's/^[^=]*=//g'`
            shift
            ;;
        --docker-registry*)
            export SWIFT_DOCKER_REGISTRY=`echo $1 | sed -e 's/^[^=]*=//g'`
            shift
            ;;
        --image-pull-secret*)
            secret=`echo $1 | sed -e 's/^[^=]*=//g'`
            export SWIFT_IMAGE_PULL_SECRET="name: '$secret'"
            shift
            ;;
        --rbac)
            export SWIFT_SERVICE_ACCOUNT=swift
            export SWIFT_ENABLE_RBAC=true
            shift
            ;;
        --run-on-master)
            export SWIFT_RUN_ON_MASTER=1
            shift
            ;;
        --uninstall)
            export SWIFT_UNINSTALL=1
            shift
            ;;
        *)
            show_help
            exit 1
            ;;
    esac
done

if [ "$SWIFT_UNINSTALL" -eq 1 ]; then
    kubectl delete deployment -l app=swift --namespace $SWIFT_NAMESPACE
    kubectl delete service -l app=swift --namespace $SWIFT_NAMESPACE
    kubectl delete secret -l app=swift --namespace $SWIFT_NAMESPACE
    # Delete RBAC objects, if --rbac flag was used.
    kubectl delete serviceaccount -l app=swift --namespace $SWIFT_NAMESPACE
    kubectl delete clusterrolebindings -l app=swift --namespace $SWIFT_NAMESPACE
    kubectl delete clusterrole -l app=swift --namespace $SWIFT_NAMESPACE

    exit 0
fi

env | sort | grep SWIFT*
echo ""

curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.7.0/hack/deploy/server.yaml | $ONESSL envsubst | kubectl apply -f -

if [ "$SWIFT_ENABLE_RBAC" = true ]; then
    kubectl create serviceaccount $SWIFT_SERVICE_ACCOUNT --namespace $SWIFT_NAMESPACE
    kubectl label serviceaccount $SWIFT_SERVICE_ACCOUNT app=swift --namespace $SWIFT_NAMESPACE
    curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.7.0/hack/deploy/rbac-list.yaml | $ONESSL envsubst | kubectl auth reconcile -f -
fi

if [ "$SWIFT_RUN_ON_MASTER" -eq 1 ]; then
    kubectl patch deploy swift -n $SWIFT_NAMESPACE \
      --patch="$(curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.7.0/hack/deploy/run-on-master.yaml)"
fi
