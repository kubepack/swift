---
title: Swift Run | Reference 
description: Swift run
menu:
  product_swift_0.6.0:
    identifier: swift-run
    name: Swift Run
    parent: reference
    weight: 14
product_name: swift
left_menu: product_swift_0.6.0 
section_menu_id: reference
---
## swift run

Run swift apis

### Synopsis


Run swift apis

```
swift run
```

### Options

```
      --api-domain string             Domain used to server swift api
      --caCertFile string             File containing CA certificate
      --certFile string               File container server TLS certificate
      --connector string              Name of connector used to connect to Tiller server. Valid values are: incluster, direct, kubeconfig, appscode
      --cors-origin-allow-subdomain   Allow CORS request from subdomains of origin
      --cors-origin-host string       Allowed CORS origin host e.g, domain[:port]
      --enable-cors                   Enable CORS support
      --enable-java-client            Set true to send SETTINGS frame from the server. Default set to false
      --keyFile string                File containing server TLS private key
      --kube-context string           Kube context used by 'kubeconfig' connection
      --plaintext-addr string         host:port used to server plaintext apis (default ":9855")
      --secure-addr string            host:port used to server secure apis (default ":50055")
      --tiller-endpoint string        Endpoint of Tiller server, eg, [scheme://]host:port
      --web-addr string               Address to listen on for web interface and telemetry. (default ":56790")
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --analytics                        Send analytical events to Google Analytics (default true)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [swift](swift.md)	 - Swift by Appscode - Ajax friendly Helm Tiller Proxy

