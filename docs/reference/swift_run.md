---
title: Swift Run
menu:
  product_swift_0.8.1:
    identifier: swift-run
    name: Swift Run
    parent: reference
product_name: swift
menu_name: product_swift_0.8.1
section_menu_id: reference
---
## swift run

Run swift apis

### Synopsis

Run swift apis

```
swift run [flags]
```

### Options

```
      --api-domain string                       Domain used for apiserver (prod: api.appscode.com
      --connector string                        Name of connector used to connect to Tiller server. Valid values are: incluster, direct, kubeconfig, appscode
      --cors-origin-allow-subdomain             Allow CORS request from subdomains of origin (default true)
      --cors-origin-host string                 Allowed CORS origin host e.g, domain[:port] (default "*")
      --enable-cors                             Enable CORS support
  -h, --help                                    help for run
      --kube-context string                     Kube context used by 'kubeconfig' connection
      --log-rpc                                 log RPC request and response
      --plaintext-addr string                   host:port used to serve http json apis (default ":9855")
      --secure-addr string                      host:port used to serve secure apis (default ":50055")
      --tiller-ca-file string                   File containing CA certificate for Tiller server
      --tiller-client-cert-file string          File container client TLS certificate for Tiller server
      --tiller-client-private-key-file string   File containing client TLS private key for Tiller server
      --tiller-endpoint string                  Endpoint of Tiller server, eg, [scheme://]host:port
      --tiller-insecure-skip-verify             Skip certificate verification for Tiller server
      --tiller-timeout duration                 Timeout used to connect to Tiller server (default 5m0s)
      --tls-ca-file string                      File containing CA certificate
      --tls-cert-file string                    File container server TLS certificate
      --tls-private-key-file string             File containing server TLS private key
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --enable-analytics                 Send analytical events to Google Analytics (default true)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [swift](/docs/reference/swift.md)	 - Swift by Appscode - Ajax friendly Helm Tiller Proxy

