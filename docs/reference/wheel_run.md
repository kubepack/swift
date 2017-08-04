## wheel run

Run wheel apis

### Synopsis


Run wheel apis

```
wheel run [flags]
```

### Options

```
      --analytics                     Send analytical events to Google Analytics (default true)
      --api-domain string             Domain used to server wheel api
      --caCertFile string             File containing CA certificate
      --certFile string               File container server TLS certificate
      --cors-origin-allow-subdomain   Allow CORS request from subdomains of origin
      --cors-origin-host string       Allowed CORS origin host e.g, domain[:port]
      --enable-cors                   Enable CORS support
      --enable-java-client            Set true to send SETTINGS frame from the server. Default set to false
  -h, --help                          help for run
      --keyFile string                File containing server TLS private key
      --plaintext-addr string         host:port used to server plaintext apis (default ":9855")
      --secure-addr string            host:port used to server secure apis (default ":50055")
      --web-addr string               Address to listen on for web interface and telemetry. (default ":5050")
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [wheel](wheel.md)	 - Wheel by Appscode - Ajax friendly Helm Tiller Proxy


