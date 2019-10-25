#go build server.go
#docker build -t server:go .
#docker run --rm -d --net=host server:go

# /usr/local/bin/wrk -t2 -c100 --latency http://localhost
Running 10s test @ http://localhost
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.20ms    1.96ms  28.52ms   86.98%
    Req/Sec    24.70k     2.71k   30.88k    75.50%
  Latency Distribution
     50%    1.68ms
     75%    2.76ms
     90%    4.05ms
     99%   10.21ms
  492784 requests in 10.03s, 60.62MB read
Requests/sec:  49133.40
Transfer/sec:      6.04MB
