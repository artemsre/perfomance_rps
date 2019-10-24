Run simple docker with 2 core and 4G ram. 

`docker run --rm -ti --net=host nginx ` 

And bomb welcome page by load:


Then try to different size files:

1K 
 dd if=/dev/urandom of=file1k.txt bs=1 count=1024
 dd if=/dev/urandom of=file10k.txt bs=1 count=10240
 dd if=/dev/urandom of=file100k.txt bs=1 count=102400
 dd if=/dev/urandom of=file1m.txt bs=1 count=1024000



Shell script 

dd if=/dev/urandom of=file1k.txt bs=1 count=1024
dd if=/dev/urandom of=file10k.txt bs=1 count=10240
dd if=/dev/urandom of=file100k.txt bs=1 count=102400
dd if=/dev/urandom of=file1m.txt bs=1 count=1024000
docker build -t perfnginx .
rm -rf file1k.txt file10k.txt file100k.txt file1m.txt
docker run --rm --net=host -d perfnginx

And after make the test: 

# /usr/local/bin/wrk -t2 -c100 --latency http://localhost
Running 10s test @ http://localhost
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.71ms    2.03ms  18.44ms   85.80%
    Req/Sec    40.41k    21.23k   78.53k    52.50%
  Latency Distribution
     50%  825.00us
     75%    2.05ms
     90%    4.82ms
     99%    8.35ms
  807784 requests in 10.10s, 124.76MB read
Requests/sec:  80015.88
Transfer/sec:     12.36MB



# /usr/local/bin/wrk -t2 -c100 --latency http://localhost/static/file1k.txt
Running 10s test @ http://localhost/static/file1k.txt
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.59ms    4.13ms 149.40ms   94.33%
    Req/Sec    22.04k     7.00k   45.58k    74.50%
  Latency Distribution
     50%    1.86ms
     75%    2.75ms
     90%    5.24ms
     99%    9.75ms
  440860 requests in 10.07s, 531.41MB read
Requests/sec:  43767.99
Transfer/sec:     52.76MB
