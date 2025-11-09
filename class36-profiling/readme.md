## class 36 - Profiling
- Profiling is measuring your program's resource usage (CPU, memory, etc.) at runtime to identify performance bottlenecks.
- `pprof` is Go's built-in profiling tool that collects and analyzes profiling data. Has a tool to find leaking go routines
- Benchmarking (`go test -bench`):
    -  Measures overall performance of specific functions
    - Answers "How fast is this code?"
    - Synthetic tests in isolation

- Profiling (`pprof`):
    - Shows WHERE time/memory is being spent in running program
    - Answers "Why is it slow?"
    - Real-world usage analysis

- Benchmarking tells you IF it's slow, profiling tells you WHY it's slow.
- Minimize other computer activity during profiling/benchmarking.

### Finding/Seeing leaking memory (goroutines)
- exercise http fetch comic 
- one goroutine is the listen socket `ListenAndServe`, starting out of `main`
- one goroutine for running `pprof`
- script that runs a bunch of curls in a row (9 queries)
- Now there are 13 goroutine (prev had run few curls) 
- lock 9 goroutine, hanging out in `netpoll` (if its a stuck in a channel is another one)
- #1 leak memory -> leak goroutine -(2 ways to leak goroutines)-> leak a socket
- issue: didn't close the body of the request to the server (go to server, read body, decode the json, never close the body!) socket gets hung up, the # of goroutines keeps going up
- how to find if goroutine leak?
   - run program and eventually dies cause it runs out of memory
   - testing run some traffic, look tool like pprof ans see how program behavies and see if the # of goroutines keeps going up even with stopped traffic.
   - to get pprof put include `_ "net/http/pprof"` `_` cause we don't use any of it, just use `http` standard lib

 ### Prometheus
 - Metrics package, give us an endpoint that shows some statistics
 - pull in packages from prometheus (github) 2: 1 define metrics and other gives a handler that we can bind that'll gives us a route to scrape the metrics from
 - in code `queries.Inc()` a counter
 - in linux, prometheus, shows open file descriptors
 - if go routines are going up and also the open fds, you are leaking sockets (if open fd are not changing but the gr are, its a bug in gr)

 ### Others

 - CPU profiling
 - Sorting, drawing animations