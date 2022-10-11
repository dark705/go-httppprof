# HTTP Pprof

Provides simple HTTP Pprof server.

## Examples

``` go
httpPprof := httppprof.NewServer(":8082")
httpPprof.Run()
defer httpPprof.Shutdown()

osSignals := make(chan os.Signal, 1)
signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

log.Printf("got signal from OS: %v. Exit...", <-osSignals)
```

After running HTTP server, you can see different profiles that go provides. In example by: 

http://localhost:8082/debug/pprof/

## Licensing

This project is licensed under the MIT License.