# link-perf
The beginnings of a link performance profiler. This is an experiment in concurrent HTTP Requests.

The struct `LinkProfiler` takes a slice of links and number of tries for each link. When executing
`Run` on the struct instance, the profiler makes concurrent requests for each link and measures various metrics
such as `GotFirstResponseByte`, `DNSStart` as per Go's [HTTP trace](https://pkg.go.dev/net/http/httptrace) module.

There are still a number of issues:

1. Probably want to use a wait group instead of manually draining channels.
2. Better error handling.
3. Fire go routines for each try?
4. App seems to stall after 3 or so links.
5. Add more metrics as per the trace interface.