## wheel run

Run wheel apis

### Synopsis


Run wheel apis

```
wheel run [flags]
```

### Options

```
      --api-port int        Port used to serve apis (default 50066)
      --caCertFile string   File containing CA certificate
      --certFile string     File container server TLS certificate
  -h, --help                help for run
      --keyFile string      File containing server TLS private key
      --pprof-port int      port used to run pprof tools (default 6060)
      --report-monitoring   Report monitoring, disabled for dev env by default
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


