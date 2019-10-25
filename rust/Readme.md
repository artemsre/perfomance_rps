# rustc server.rs 
# docker build -t server:rust .
# docker run --rm -d --net=host server:rust

# /usr/local/bin/wrk -t2 -c100 --latency http://localhost
Running 10s test @ http://localhost
  2 threads and 100 connections
    Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    20.93ms   18.81ms 420.41ms   97.24%
    Req/Sec     1.70k   327.59     2.36k    85.16%
  Latency Distribution
     50%   18.22ms
     75%   20.19ms
     90%   22.64ms
     99%  109.98ms
  21758 requests in 6.44s, 1.49MB read
  Socket errors: connect 0, read 21758, write 0, timeout 0
Requests/sec:   3378.11
Transfer/sec:    237.61KB
