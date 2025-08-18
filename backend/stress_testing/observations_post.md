Here when I do pressure testing on the most used post request that is sending food to users to chef
```
ab -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImdveWFsLmFyeWFuQGdtYWlsLmNvbSIsImV4cCI6MTc1NTUyODMwMiwicm9sZSI6ImN1c3RvbWVyIn0.DD5pD47SZeXoYICcvG8crdkUML3MbrfYuer9tYlxddg" -c 1000 -n 1000000 -p data_send.json -T "application/json" -l http://127.0.0.1:8000/customer/food_items_added
```

The result that I get is
```
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100000 requests
Completed 200000 requests
Completed 300000 requests
Completed 400000 requests
Completed 500000 requests
Completed 600000 requests
Completed 700000 requests
Completed 800000 requests
Completed 900000 requests
Completed 1000000 requests
Finished 1000000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8000

Document Path:          /customer/food_items_added
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   458.158 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      279000000 bytes
Total body sent:        497000000
HTML transferred:       40000000 bytes
Requests per second:    2182.65 [#/sec] (mean)
Time per request:       458.158 [ms] (mean)
Time per request:       0.458 [ms] (mean, across all concurrent requests)
Transfer rate:          594.69 [Kbytes/sec] received
                        1059.35 kb/s sent
                        1654.04 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   2.7      0     188
Processing:    55  457 134.1    442    1701
Waiting:       33  457 134.1    442    1701
Total:         55  458 134.1    443    1702

Percentage of the requests served within a certain time (ms)
  50%    443
  66%    500
  75%    538
  80%    564
  90%    635
  95%    699
  98%    776
  99%    834
 100%   1702 (longest request)
 ```