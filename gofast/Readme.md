#go get "github.com/valyala/fasthttp"
#go build server.go
#docker build -t server:fastgo .
#docker run --rm -d --net=host server:fastgo

# /usr/local/bin/wrk -t2 -c100 --latency http://localhost/static/file1m.txt
Running 10s test @ http://localhost/static/file1m.txt
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.23ms    1.81ms  70.61ms   93.64%
    Req/Sec    45.55k     7.30k   64.91k    75.00%
  Latency Distribution
     50%    0.86ms
     75%    1.11ms
     90%    2.00ms
     99%    7.65ms
  908240 requests in 10.04s, 129.06MB read
Requests/sec:  90487.64
Transfer/sec:     12.86MB

