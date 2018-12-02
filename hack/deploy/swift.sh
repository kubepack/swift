#!/bin/bash
set -eou pipefail

echo "checking kubeconfig context"
kubectl config current-context || {
  echo "Set a context (kubectl use-context <context>) out of the following:"
  echo
  kubectl config get-contexts
  exit 1
}
echo ""

# http://redsymbol.net/articles/bash-exit-traps/
function cleanup() {
  rm -rf $ONESSL ca.crt ca.key server.crt server.key
}
trap cleanup EXIT

# ref: https://github.com/appscodelabs/libbuild/blob/master/common/lib.sh#L55
inside_git_repo() {
  git rev-parse --is-inside-work-tree >/dev/null 2>&1
  inside_git=$?
  if [ "$inside_git" -ne 0 ]; then
    echo "Not inside a git repository"
    exit 1
  fi
}

detect_tag() {
  inside_git_repo

  # http://stackoverflow.com/a/1404862/3476121
  git_tag=$(git describe --exact-match --abbrev=0 2>/dev/null || echo '')

  commit_hash=$(git rev-parse --verify HEAD)
  git_branch=$(git rev-parse --abbrev-ref HEAD)
  commit_timestamp=$(git show -s --format=%ct)

  if [ "$git_tag" != '' ]; then
    TAG=$git_tag
    TAG_STRATEGY='git_tag'
  elif [ "$git_branch" != 'master' ] && [ "$git_branch" != 'HEAD' ] && [[ "$git_branch" != release-* ]]; then
    TAG=$git_branch
    TAG_STRATEGY='git_branch'
  else
    hash_ver=$(git describe --tags --always --dirty)
    TAG="${hash_ver}"
    TAG_STRATEGY='commit_hash'
  fi

  export TAG
  export TAG_STRATEGY
  export git_tag
  export git_branch
  export commit_hash
  export commit_timestamp
}

onessl_found() {
  # https://stackoverflow.com/a/677212/244009
  if [ -x "$(command -v onessl)" ]; then
    onessl wait-until-has -h >/dev/null 2>&1 || {
      # old version of onessl found
      echo "Found outdated onessl"
      return 1
    }
    export ONESSL=onessl
    return 0
  fi
  return 1
}

onessl_found || {
  echo "Downloading onessl ..."
  # ref: https://stackoverflow.com/a/27776822/244009
  case "$(uname -s)" in
    Darwin)
      curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.9.0/onessl-darwin-amd64
      chmod +x onessl
      export ONESSL=./onessl
      ;;

    Linux)
      curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.9.0/onessl-linux-amd64
      chmod +x onessl
      export ONESSL=./onessl
      ;;

    CYGWIN* | MINGW32* | MSYS*)
      curl -fsSL -o onessl.exe https://github.com/kubepack/onessl/releases/download/0.9.0/onessl-windows-amd64.exe
      chmod +x onessl.exe
      export ONESSL=./onessl.exe
      ;;
    *)
      echo 'other OS'
      ;;
  esac
}

# ref: https://stackoverflow.com/a/7069755/244009
# ref: https://jonalmeida.com/posts/2013/05/26/different-ways-to-implement-flags-in-bash/
# ref: http://tldp.org/LDP/abs/html/comparison-ops.html

export SWIFT_NAMESPACE=kube-system
export SWIFT_SERVICE_ACCOUNT=swift
export SWIFT_ENABLE_RBAC=true
export SWIFT_RUN_ON_MASTER=0
export SWIFT_DOCKER_REGISTRY=appscode
export SWIFT_SERVER_TAG=0.9.0
export SWIFT_IMAGE_PULL_SECRET=
export SWIFT_IMAGE_PULL_POLICY=IfNotPresent
export SWIFT_ENABLE_ANALYTICS=true
export SWIFT_UNINSTALL=0

export APPSCODE_ENV=${APPSCODE_ENV:-prod}
export SCRIPT_LOCATION="curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.9.0/"
if [ "$APPSCODE_ENV" = "dev" ]; then
  detect_tag
  export SCRIPT_LOCATION="cat "
  export SWIFT_SERVER_TAG=$TAG
  export SWIFT_IMAGE_PULL_POLICY=Always
fi

show_help() {
  echo "swift.sh - install Ajax friendly Helm Tiller Proxy"
  echo " "
  echo "swift.sh [options]"
  echo " "
  echo "options:"
  echo "-h, --help                         show brief help"
  echo "-n, --namespace=NAMESPACE          specify namespace (default: kube-system)"
  echo "    --rbac                         create RBAC roles and bindings (default: true)"
  echo "    --docker-registry              docker registry used to pull swift images (default: appscode)"
  echo "    --image-pull-secret            name of secret used to pull swift docker images"
  echo "    --run-on-master                run swift operator on master"
  echo "    --enable-analytics             send usage events to Google Analytics (default: true)"
  echo "    --uninstall                    uninstall swift"
}

while test $# -gt 0; do
  case "$1" in
    -h | --help)
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
      export SWIFT_NAMESPACE=$(echo $1 | sed -e 's/^[^=]*=//g')
      shift
      ;;
    --docker-registry*)
      export SWIFT_DOCKER_REGISTRY=$(echo $1 | sed -e 's/^[^=]*=//g')
      shift
      ;;
    --image-pull-secret*)
      secret=$(echo $1 | sed -e 's/^[^=]*=//g')
      export SWIFT_IMAGE_PULL_SECRET="name: '$secret'"
      shift
      ;;
    --rbac*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export SWIFT_SERVICE_ACCOUNT=default
        export SWIFT_ENABLE_RBAC=false
      fi
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

${SCRIPT_LOCATION}hack/deploy/server.yaml | $ONESSL envsubst | kubectl apply -f -

if [ "$SWIFT_ENABLE_RBAC" = true ]; then
  ${SCRIPT_LOCATION}hack/deploy/service-account.yaml | $ONESSL envsubst | kubectl apply -f -
  ${SCRIPT_LOCATION}hack/deploy/rbac-list.yaml | $ONESSL envsubst | kubectl auth reconcile -f -
fi

if [ "$SWIFT_RUN_ON_MASTER" -eq 1 ]; then
  kubectl patch deploy swift -n $SWIFT_NAMESPACE \
    --patch="$(${SCRIPT_LOCATION}hack/deploy/run-on-master.yaml)"
fi

echo
echo "Successfully installed Swift in $SWIFT_NAMESPACE namespace!"
