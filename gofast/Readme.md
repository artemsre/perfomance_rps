#go get "github.com/valyala/fasthttp"
#go build server.go
#docker build -t server:fastgo .
#docker run --rm -d --net=host server:fastgo

# /usr/local/bin/wrk -t2 -c100 --latency http://localhost
Running 10s test @ http://localhost
  2 threads and 100 connections
    Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.76ms    2.27ms  32.93ms   86.36%
    Req/Sec    42.63k    18.00k   82.84k    64.52%
  Latency Distribution
     50%  762.00us
     75%    2.13ms
     90%    5.03ms
     99%    9.49ms
  264300 requests in 3.14s, 40.82MB read
Requests/sec:  84139.34
Transfer/sec:     12.16MB
