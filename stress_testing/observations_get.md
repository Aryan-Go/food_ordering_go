When I run the following command on my menu ordering route then this is the result I get

```
ab -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImdveWFsLmFyeWFuQGdtYWlsLmNvbSIsImV4cCI6MTc1NTYzMDY4Miwicm9sZSI6ImN1c3RvbWVyIn0.vingHfG0upimFnqBpZtOv__rJpKdXrkvAL7z3KCIPPc" -c 1000 -n 100000 http://127.0.0.1:8000/customer/menu_show
```

This is the reult that I get

```
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /customer/menu_show
Document Length:        2365 bytes

Concurrency Level:      1000
Time taken for tests:   3.955 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      258400000 bytes
HTML transferred:       236500000 bytes
Requests per second:    25286.58 [#/sec] (mean)
Time per request:       39.547 [ms] (mean)
Time per request:       0.040 [ms] (mean, across all concurrent requests)
Transfer rate:          63809.10 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   15  51.8      4    2014
Processing:     0    6   5.8      5      95
Waiting:        0    6   5.6      5      95
Total:          0   21  52.3      9    2054

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     11
  75%     12
  80%     14
  90%     40
  95%     71
  98%    131
  99%    220
 100%   2054 (longest request)

```

Hence it passed here

For Orders fetching for chef

If I run the command
```
ab -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoZWZAZ21haWwuY29tIiwiZXhwIjoxNzU1NTI1MzQ0LCJyb2xlIjoiY2hlZiJ9.MdjpqtmVHNVh2cgS2JTu5EvGfiHmChjowXQuh4GEzkM" -c 100 -n 1000 -s 120 http://127.0.0.1:8000/chef/render_order         
```
Then the result I get is
```
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /chef/render_order
Document Length:        99 bytes

Concurrency Level:      100
Time taken for tests:   0.146 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      338000 bytes
HTML transferred:       99000 bytes
Requests per second:    6845.14 [#/sec] (mean)
Time per request:       14.609 [ms] (mean)
Time per request:       0.146 [ms] (mean, across all concurrent requests)
Transfer rate:          2259.43 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       3
Processing:     3   13   7.8     11      50
Waiting:        3   13   7.8     11      50
Total:          3   14   8.3     11      52

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     13
  75%     15
  80%     16
  90%     25
  95%     34
  98%     40
  99%     45
 100%     52 (longest request)
```

When I run the same for 1000 concurrency and 1000000 requests then the results are as follows

```
Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /chef/render_order
Document Length:        99 bytes

Concurrency Level:      1000
Time taken for tests:   20.668 seconds
Complete requests:      15
Failed requests:        0
Total transferred:      5070 bytes
HTML transferred:       1485 bytes
Requests per second:    0.73 [#/sec] (mean)
Time per request:       1377851.333 [ms] (mean)
Time per request:       1377.851 [ms] (mean, across all concurrent requests)
Transfer rate:          0.24 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   15   8.6     16      33
Processing:    24   43   9.0     47      49
Waiting:        3   41  12.8     46      49
Total:         25   57  11.5     59      71

Percentage of the requests served within a certain time (ms)
  50%     58
  66%     60
  75%     68
  80%     69
  90%     70
  95%     71
  98%     71
  99%     71
 100%     71 (longest request)
 ```

 