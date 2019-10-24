
# docker build -t server:python .
# docker run --rm -d --net=host server:python

# /usr/local/bin/wrk -t2 -c100 --latency http://localhost/
Running 10s test @ http://localhost/
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.14ms   80.13ms   1.61s    98.33%
    Req/Sec     1.55k     0.93k    3.52k    70.89%
  Latency Distribution
     50%  683.00us
     75%    0.86ms
     90%    1.10ms
     99%  313.17ms
  12420 requests in 10.10s, 1.23MB read
  Socket errors: connect 0, read 12492, write 0, timeout 4
Requests/sec:   1229.58
Transfer/sec:    124.88KB

