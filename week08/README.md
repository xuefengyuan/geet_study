## 第一题

使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

### 一、set/get性能测试

并发请求的Client 20个

发送总请求数为100000

value的大小为 10字节

key的范围为10000内随机

```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -a redis -c 20 -n 1000000 -d 10b -r 100000 -t set,get
====== SET ======                                                     
  1000000 requests completed in 5.29 seconds
  20 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.031 milliseconds (cumulative count 4)
50.000% <= 0.055 milliseconds (cumulative count 528757)
75.000% <= 0.063 milliseconds (cumulative count 780978)
87.500% <= 0.079 milliseconds (cumulative count 923679)
93.750% <= 0.087 milliseconds (cumulative count 957783)
96.875% <= 0.095 milliseconds (cumulative count 972334)
98.438% <= 0.111 milliseconds (cumulative count 985566)
99.219% <= 0.143 milliseconds (cumulative count 992241)
99.609% <= 0.199 milliseconds (cumulative count 996298)
99.805% <= 0.255 milliseconds (cumulative count 998196)
99.902% <= 0.327 milliseconds (cumulative count 999071)
99.951% <= 0.415 milliseconds (cumulative count 999520)
99.976% <= 0.519 milliseconds (cumulative count 999764)
99.988% <= 0.655 milliseconds (cumulative count 999880)
99.994% <= 0.951 milliseconds (cumulative count 999939)
99.997% <= 1.639 milliseconds (cumulative count 999970)
99.998% <= 1.687 milliseconds (cumulative count 999987)
99.999% <= 1.711 milliseconds (cumulative count 999994)
100.000% <= 1.727 milliseconds (cumulative count 999997)
100.000% <= 1.767 milliseconds (cumulative count 999999)
100.000% <= 1.815 milliseconds (cumulative count 1000000)
100.000% <= 1.815 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
98.078% <= 0.103 milliseconds (cumulative count 980782)
99.668% <= 0.207 milliseconds (cumulative count 996678)
99.886% <= 0.303 milliseconds (cumulative count 998856)
99.948% <= 0.407 milliseconds (cumulative count 999478)
99.975% <= 0.503 milliseconds (cumulative count 999753)
99.985% <= 0.607 milliseconds (cumulative count 999854)
99.989% <= 0.703 milliseconds (cumulative count 999893)
99.991% <= 0.807 milliseconds (cumulative count 999911)
99.993% <= 0.903 milliseconds (cumulative count 999929)
99.994% <= 1.007 milliseconds (cumulative count 999940)
99.995% <= 1.303 milliseconds (cumulative count 999948)
99.996% <= 1.407 milliseconds (cumulative count 999960)
99.996% <= 1.607 milliseconds (cumulative count 999961)
99.999% <= 1.703 milliseconds (cumulative count 999991)
100.000% <= 1.807 milliseconds (cumulative count 999999)
100.000% <= 1.903 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 189071.66 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.061     0.024     0.055     0.087     0.135     1.815
====== GET ======                                                     
  1000000 requests completed in 5.29 seconds
  20 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 2)
50.000% <= 0.055 milliseconds (cumulative count 592162)
75.000% <= 0.063 milliseconds (cumulative count 754070)
87.500% <= 0.079 milliseconds (cumulative count 931006)
93.750% <= 0.087 milliseconds (cumulative count 959465)
96.875% <= 0.095 milliseconds (cumulative count 975452)
98.438% <= 0.111 milliseconds (cumulative count 988201)
99.219% <= 0.135 milliseconds (cumulative count 992586)
99.609% <= 0.191 milliseconds (cumulative count 996262)
99.805% <= 0.247 milliseconds (cumulative count 998145)
99.902% <= 0.327 milliseconds (cumulative count 999039)
99.951% <= 0.415 milliseconds (cumulative count 999513)
99.976% <= 0.471 milliseconds (cumulative count 999772)
99.988% <= 0.591 milliseconds (cumulative count 999881)
99.994% <= 0.855 milliseconds (cumulative count 999939)
99.997% <= 1.007 milliseconds (cumulative count 999970)
99.998% <= 1.367 milliseconds (cumulative count 999985)
99.999% <= 1.415 milliseconds (cumulative count 999999)
100.000% <= 1.423 milliseconds (cumulative count 1000000)
100.000% <= 1.423 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
98.381% <= 0.103 milliseconds (cumulative count 983812)
99.700% <= 0.207 milliseconds (cumulative count 997002)
99.886% <= 0.303 milliseconds (cumulative count 998858)
99.945% <= 0.407 milliseconds (cumulative count 999453)
99.980% <= 0.503 milliseconds (cumulative count 999796)
99.988% <= 0.607 milliseconds (cumulative count 999881)
99.993% <= 0.703 milliseconds (cumulative count 999927)
99.994% <= 0.807 milliseconds (cumulative count 999938)
99.995% <= 0.903 milliseconds (cumulative count 999948)
99.997% <= 1.007 milliseconds (cumulative count 999970)
99.998% <= 1.207 milliseconds (cumulative count 999980)
99.999% <= 1.407 milliseconds (cumulative count 999992)
100.000% <= 1.503 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 189178.97 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.060     0.016     0.055     0.087     0.119     1.423

```

并发请求的Client 20个

发送总请求数为100000

value的大小为 20字节

key的范围为10000内随机

```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -a redis -c 20 -n 1000000 -d 20b -r 100000 -t set,get
====== SET ======                                                     
  1000000 requests completed in 5.81 seconds
  20 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 1)
50.000% <= 0.063 milliseconds (cumulative count 608408)
75.000% <= 0.079 milliseconds (cumulative count 830922)
87.500% <= 0.087 milliseconds (cumulative count 908296)
93.750% <= 0.095 milliseconds (cumulative count 944430)
96.875% <= 0.111 milliseconds (cumulative count 979775)
98.438% <= 0.119 milliseconds (cumulative count 985105)
99.219% <= 0.151 milliseconds (cumulative count 992512)
99.609% <= 0.199 milliseconds (cumulative count 996232)
99.805% <= 0.263 milliseconds (cumulative count 998171)
99.902% <= 0.335 milliseconds (cumulative count 999051)
99.951% <= 0.415 milliseconds (cumulative count 999548)
99.976% <= 0.455 milliseconds (cumulative count 999774)
99.988% <= 0.503 milliseconds (cumulative count 999884)
99.994% <= 0.575 milliseconds (cumulative count 999944)
99.997% <= 1.255 milliseconds (cumulative count 999970)
99.998% <= 2.679 milliseconds (cumulative count 999985)
99.999% <= 2.735 milliseconds (cumulative count 999994)
100.000% <= 2.759 milliseconds (cumulative count 999997)
100.000% <= 2.775 milliseconds (cumulative count 999999)
100.000% <= 2.807 milliseconds (cumulative count 1000000)
100.000% <= 2.807 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
96.874% <= 0.103 milliseconds (cumulative count 968742)
99.658% <= 0.207 milliseconds (cumulative count 996581)
99.877% <= 0.303 milliseconds (cumulative count 998769)
99.950% <= 0.407 milliseconds (cumulative count 999499)
99.988% <= 0.503 milliseconds (cumulative count 999884)
99.995% <= 0.607 milliseconds (cumulative count 999946)
99.996% <= 0.703 milliseconds (cumulative count 999958)
99.997% <= 1.007 milliseconds (cumulative count 999969)
99.998% <= 1.303 milliseconds (cumulative count 999980)
100.000% <= 3.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 172117.05 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.067     0.016     0.063     0.103     0.135     2.807
====== GET ======                                                     
  1000000 requests completed in 5.57 seconds
  20 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 1)
50.000% <= 0.055 milliseconds (cumulative count 504917)
75.000% <= 0.071 milliseconds (cumulative count 751892)
87.500% <= 0.087 milliseconds (cumulative count 923039)
93.750% <= 0.095 milliseconds (cumulative count 955220)
96.875% <= 0.103 milliseconds (cumulative count 974633)
98.438% <= 0.119 milliseconds (cumulative count 988067)
99.219% <= 0.143 milliseconds (cumulative count 992581)
99.609% <= 0.199 milliseconds (cumulative count 996424)
99.805% <= 0.263 milliseconds (cumulative count 998182)
99.902% <= 0.327 milliseconds (cumulative count 999066)
99.951% <= 0.415 milliseconds (cumulative count 999515)
99.976% <= 0.487 milliseconds (cumulative count 999772)
99.988% <= 0.599 milliseconds (cumulative count 999878)
99.994% <= 0.831 milliseconds (cumulative count 999946)
99.997% <= 1.303 milliseconds (cumulative count 999973)
99.998% <= 1.503 milliseconds (cumulative count 999985)
99.999% <= 1.591 milliseconds (cumulative count 1000000)
100.000% <= 1.591 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
97.463% <= 0.103 milliseconds (cumulative count 974633)
99.674% <= 0.207 milliseconds (cumulative count 996744)
99.885% <= 0.303 milliseconds (cumulative count 998853)
99.945% <= 0.407 milliseconds (cumulative count 999449)
99.978% <= 0.503 milliseconds (cumulative count 999783)
99.988% <= 0.607 milliseconds (cumulative count 999883)
99.991% <= 0.703 milliseconds (cumulative count 999911)
99.991% <= 0.807 milliseconds (cumulative count 999912)
99.995% <= 0.903 milliseconds (cumulative count 999946)
99.996% <= 1.007 milliseconds (cumulative count 999956)
99.997% <= 1.103 milliseconds (cumulative count 999967)
99.997% <= 1.303 milliseconds (cumulative count 999973)
99.998% <= 1.407 milliseconds (cumulative count 999978)
99.999% <= 1.503 milliseconds (cumulative count 999985)
100.000% <= 1.607 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 179372.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.064     0.016     0.055     0.095     0.127     1.591


```

并发请求的Client 20个

发送总请求数为100000

value的大小为 50字节

key的范围为10000内随机

```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -a redis -c 20 -n 1000000 -d 50b -r 100000 -t set,get
====== SET ======                                                     
  1000000 requests completed in 5.96 seconds
  20 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 1)
50.000% <= 0.063 milliseconds (cumulative count 550324)
75.000% <= 0.079 milliseconds (cumulative count 781797)
87.500% <= 0.087 milliseconds (cumulative count 883584)
93.750% <= 0.103 milliseconds (cumulative count 956277)
96.875% <= 0.111 milliseconds (cumulative count 970651)
98.438% <= 0.135 milliseconds (cumulative count 985731)
99.219% <= 0.175 milliseconds (cumulative count 992916)
99.609% <= 0.223 milliseconds (cumulative count 996216)
99.805% <= 0.279 milliseconds (cumulative count 998150)
99.902% <= 0.359 milliseconds (cumulative count 999051)
99.951% <= 0.447 milliseconds (cumulative count 999515)
99.976% <= 0.503 milliseconds (cumulative count 999763)
99.988% <= 0.607 milliseconds (cumulative count 999881)
99.994% <= 0.879 milliseconds (cumulative count 999939)
99.997% <= 1.463 milliseconds (cumulative count 999972)
99.998% <= 2.535 milliseconds (cumulative count 999985)
99.999% <= 2.583 milliseconds (cumulative count 999993)
100.000% <= 2.607 milliseconds (cumulative count 999997)
100.000% <= 2.623 milliseconds (cumulative count 999999)
100.000% <= 2.655 milliseconds (cumulative count 1000000)
100.000% <= 2.655 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
95.628% <= 0.103 milliseconds (cumulative count 956277)
99.541% <= 0.207 milliseconds (cumulative count 995412)
99.856% <= 0.303 milliseconds (cumulative count 998558)
99.928% <= 0.407 milliseconds (cumulative count 999282)
99.976% <= 0.503 milliseconds (cumulative count 999763)
99.988% <= 0.607 milliseconds (cumulative count 999881)
99.992% <= 0.703 milliseconds (cumulative count 999920)
99.993% <= 0.807 milliseconds (cumulative count 999929)
99.994% <= 0.903 milliseconds (cumulative count 999939)
99.994% <= 1.007 milliseconds (cumulative count 999940)
99.995% <= 1.207 milliseconds (cumulative count 999945)
99.995% <= 1.303 milliseconds (cumulative count 999953)
99.996% <= 1.407 milliseconds (cumulative count 999961)
99.998% <= 1.503 milliseconds (cumulative count 999977)
99.998% <= 1.607 milliseconds (cumulative count 999980)
100.000% <= 3.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 167672.70 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.070     0.016     0.063     0.103     0.159     2.655
====== GET ======                                                     
  1000000 requests completed in 6.21 seconds
  20 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 2)
50.000% <= 0.071 milliseconds (cumulative count 538103)
75.000% <= 0.087 milliseconds (cumulative count 834775)
87.500% <= 0.095 milliseconds (cumulative count 902745)
93.750% <= 0.103 milliseconds (cumulative count 946689)
96.875% <= 0.111 milliseconds (cumulative count 969557)
98.438% <= 0.127 milliseconds (cumulative count 986148)
99.219% <= 0.159 milliseconds (cumulative count 992939)
99.609% <= 0.215 milliseconds (cumulative count 996312)
99.805% <= 0.279 milliseconds (cumulative count 998137)
99.902% <= 0.359 milliseconds (cumulative count 999024)
99.951% <= 0.439 milliseconds (cumulative count 999529)
99.976% <= 0.479 milliseconds (cumulative count 999769)
99.988% <= 0.567 milliseconds (cumulative count 999878)
99.994% <= 0.671 milliseconds (cumulative count 999939)
99.997% <= 0.767 milliseconds (cumulative count 999978)
99.998% <= 0.807 milliseconds (cumulative count 999988)
99.999% <= 1.207 milliseconds (cumulative count 999998)
100.000% <= 1.215 milliseconds (cumulative count 1000000)
100.000% <= 1.215 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
94.669% <= 0.103 milliseconds (cumulative count 946689)
99.599% <= 0.207 milliseconds (cumulative count 995987)
99.854% <= 0.303 milliseconds (cumulative count 998537)
99.935% <= 0.407 milliseconds (cumulative count 999348)
99.983% <= 0.503 milliseconds (cumulative count 999834)
99.990% <= 0.607 milliseconds (cumulative count 999901)
99.995% <= 0.703 milliseconds (cumulative count 999949)
99.999% <= 0.807 milliseconds (cumulative count 999988)
99.999% <= 0.903 milliseconds (cumulative count 999990)
100.000% <= 1.207 milliseconds (cumulative count 999998)
100.000% <= 1.303 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 161004.67 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.071     0.016     0.071     0.111     0.143     1.215

```

并发请求的Client 20个

发送总请求数为100000

value的大小为 100 字节

key的范围为10000内随机

```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -a redis -c 20 -n 1000000 -d 100b -r 100000 -t set,get
====== SET ======                                                     
  1000000 requests completed in 5.29 seconds
  20 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 1)
50.000% <= 0.063 milliseconds (cumulative count 689429)
75.000% <= 0.071 milliseconds (cumulative count 828360)
87.500% <= 0.079 milliseconds (cumulative count 908323)
93.750% <= 0.087 milliseconds (cumulative count 952874)
96.875% <= 0.095 milliseconds (cumulative count 970816)
98.438% <= 0.111 milliseconds (cumulative count 985035)
99.219% <= 0.143 milliseconds (cumulative count 992433)
99.609% <= 0.191 milliseconds (cumulative count 996277)
99.805% <= 0.239 milliseconds (cumulative count 998070)
99.902% <= 0.311 milliseconds (cumulative count 999055)
99.951% <= 0.375 milliseconds (cumulative count 999520)
99.976% <= 0.423 milliseconds (cumulative count 999762)
99.988% <= 0.447 milliseconds (cumulative count 999907)
99.994% <= 0.463 milliseconds (cumulative count 999940)
99.997% <= 0.535 milliseconds (cumulative count 999970)
99.998% <= 0.607 milliseconds (cumulative count 999986)
99.999% <= 0.679 milliseconds (cumulative count 999999)
100.000% <= 0.687 milliseconds (cumulative count 1000000)
100.000% <= 0.687 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
98.002% <= 0.103 milliseconds (cumulative count 980017)
99.708% <= 0.207 milliseconds (cumulative count 997077)
99.898% <= 0.303 milliseconds (cumulative count 998976)
99.965% <= 0.407 milliseconds (cumulative count 999648)
99.996% <= 0.503 milliseconds (cumulative count 999959)
99.999% <= 0.607 milliseconds (cumulative count 999986)
100.000% <= 0.703 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 189178.97 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.063     0.016     0.063     0.087     0.135     0.687
====== GET ======                                                     
  1000000 requests completed in 5.19 seconds
  20 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 1)
50.000% <= 0.055 milliseconds (cumulative count 659441)
75.000% <= 0.063 milliseconds (cumulative count 842615)
87.500% <= 0.071 milliseconds (cumulative count 905837)
93.750% <= 0.079 milliseconds (cumulative count 961869)
96.875% <= 0.087 milliseconds (cumulative count 976107)
98.438% <= 0.103 milliseconds (cumulative count 988019)
99.219% <= 0.135 milliseconds (cumulative count 992911)
99.609% <= 0.183 milliseconds (cumulative count 996236)
99.805% <= 0.239 milliseconds (cumulative count 998175)
99.902% <= 0.319 milliseconds (cumulative count 999043)
99.951% <= 0.423 milliseconds (cumulative count 999518)
99.976% <= 0.519 milliseconds (cumulative count 999791)
99.988% <= 0.583 milliseconds (cumulative count 999891)
99.994% <= 0.743 milliseconds (cumulative count 999943)
99.997% <= 1.239 milliseconds (cumulative count 999970)
99.998% <= 1.351 milliseconds (cumulative count 999985)
99.999% <= 1.663 milliseconds (cumulative count 999993)
100.000% <= 1.695 milliseconds (cumulative count 999997)
100.000% <= 1.727 milliseconds (cumulative count 1000000)
100.000% <= 1.727 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
98.802% <= 0.103 milliseconds (cumulative count 988019)
99.723% <= 0.207 milliseconds (cumulative count 997230)
99.891% <= 0.303 milliseconds (cumulative count 998910)
99.944% <= 0.407 milliseconds (cumulative count 999437)
99.975% <= 0.503 milliseconds (cumulative count 999747)
99.992% <= 0.607 milliseconds (cumulative count 999915)
99.992% <= 0.703 milliseconds (cumulative count 999923)
99.994% <= 0.807 milliseconds (cumulative count 999943)
99.996% <= 0.903 milliseconds (cumulative count 999955)
99.997% <= 1.103 milliseconds (cumulative count 999968)
99.998% <= 1.303 milliseconds (cumulative count 999981)
99.999% <= 1.407 milliseconds (cumulative count 999989)
100.000% <= 1.703 milliseconds (cumulative count 999997)
100.000% <= 1.807 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 192752.50 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.058     0.016     0.055     0.079     0.111     1.727
```

并发请求的Client 20个

发送总请求数为100000

value的大小为 200 字节

key的范围为10000内随机

```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -a redis -c 20 -n 1000000 -d 200b -r 100000 -t set,get
====== SET ======                                                     
  1000000 requests completed in 5.40 seconds
  20 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.031 milliseconds (cumulative count 2)
50.000% <= 0.063 milliseconds (cumulative count 611084)
75.000% <= 0.071 milliseconds (cumulative count 771653)
87.500% <= 0.079 milliseconds (cumulative count 875084)
93.750% <= 0.095 milliseconds (cumulative count 959835)
96.875% <= 0.103 milliseconds (cumulative count 973337)
98.438% <= 0.119 milliseconds (cumulative count 984538)
99.219% <= 0.159 milliseconds (cumulative count 992551)
99.609% <= 0.207 milliseconds (cumulative count 996258)
99.805% <= 0.263 milliseconds (cumulative count 998152)
99.902% <= 0.327 milliseconds (cumulative count 999050)
99.951% <= 0.407 milliseconds (cumulative count 999566)
99.976% <= 0.439 milliseconds (cumulative count 999782)
99.988% <= 0.471 milliseconds (cumulative count 999880)
99.994% <= 0.591 milliseconds (cumulative count 999943)
99.997% <= 0.647 milliseconds (cumulative count 999973)
99.998% <= 1.199 milliseconds (cumulative count 999986)
99.999% <= 1.279 milliseconds (cumulative count 999994)
100.000% <= 1.303 milliseconds (cumulative count 999997)
100.000% <= 1.319 milliseconds (cumulative count 999999)
100.000% <= 1.359 milliseconds (cumulative count 1000000)
100.000% <= 1.359 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
97.334% <= 0.103 milliseconds (cumulative count 973337)
99.626% <= 0.207 milliseconds (cumulative count 996258)
99.883% <= 0.303 milliseconds (cumulative count 998827)
99.957% <= 0.407 milliseconds (cumulative count 999566)
99.991% <= 0.503 milliseconds (cumulative count 999906)
99.996% <= 0.607 milliseconds (cumulative count 999963)
99.998% <= 0.703 milliseconds (cumulative count 999980)
99.999% <= 1.207 milliseconds (cumulative count 999987)
100.000% <= 1.303 milliseconds (cumulative count 999997)
100.000% <= 1.407 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 185048.12 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.065     0.024     0.063     0.095     0.143     1.359
====== GET ======                                                     
  1000000 requests completed in 5.53 seconds
  20 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 2)
50.000% <= 0.055 milliseconds (cumulative count 555689)
75.000% <= 0.071 milliseconds (cumulative count 792664)
87.500% <= 0.079 milliseconds (cumulative count 892827)
93.750% <= 0.095 milliseconds (cumulative count 957808)
96.875% <= 0.103 milliseconds (cumulative count 973399)
98.438% <= 0.119 milliseconds (cumulative count 986330)
99.219% <= 0.151 milliseconds (cumulative count 992557)
99.609% <= 0.199 milliseconds (cumulative count 996155)
99.805% <= 0.255 milliseconds (cumulative count 998241)
99.902% <= 0.311 milliseconds (cumulative count 999035)
99.951% <= 0.399 milliseconds (cumulative count 999519)
99.976% <= 0.455 milliseconds (cumulative count 999768)
99.988% <= 0.519 milliseconds (cumulative count 999888)
99.994% <= 0.551 milliseconds (cumulative count 999943)
99.997% <= 0.639 milliseconds (cumulative count 999978)
99.998% <= 1.391 milliseconds (cumulative count 999989)
99.999% <= 1.599 milliseconds (cumulative count 1000000)
100.000% <= 1.599 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
97.340% <= 0.103 milliseconds (cumulative count 973399)
99.652% <= 0.207 milliseconds (cumulative count 996518)
99.895% <= 0.303 milliseconds (cumulative count 998952)
99.955% <= 0.407 milliseconds (cumulative count 999550)
99.985% <= 0.503 milliseconds (cumulative count 999849)
99.996% <= 0.607 milliseconds (cumulative count 999965)
99.998% <= 0.703 milliseconds (cumulative count 999979)
99.999% <= 1.407 milliseconds (cumulative count 999989)
100.000% <= 1.607 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 180962.72 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.062     0.016     0.055     0.095     0.135     1.599

```

并发请求的Client 20 个

发送总请求数为100000

value的大小为 1K字节

key的范围为10000内随机

```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -a redis -c 20 -n 1000000 -d 1k -r 100000 -t set,get
====== SET ======                                                     
  1000000 requests completed in 5.23 seconds
  20 parallel clients
  1 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 2)
50.000% <= 0.055 milliseconds (cumulative count 565495)
75.000% <= 0.063 milliseconds (cumulative count 793820)
87.500% <= 0.079 milliseconds (cumulative count 928869)
93.750% <= 0.087 milliseconds (cumulative count 958363)
96.875% <= 0.095 milliseconds (cumulative count 973784)
98.438% <= 0.111 milliseconds (cumulative count 986788)
99.219% <= 0.143 milliseconds (cumulative count 992887)
99.609% <= 0.199 milliseconds (cumulative count 996366)
99.805% <= 0.255 milliseconds (cumulative count 998059)
99.902% <= 0.319 milliseconds (cumulative count 999034)
99.951% <= 0.407 milliseconds (cumulative count 999514)
99.976% <= 0.455 milliseconds (cumulative count 999791)
99.988% <= 0.495 milliseconds (cumulative count 999880)
99.994% <= 0.671 milliseconds (cumulative count 999940)
99.997% <= 0.895 milliseconds (cumulative count 999970)
99.998% <= 2.415 milliseconds (cumulative count 999985)
99.999% <= 2.519 milliseconds (cumulative count 999993)
100.000% <= 2.575 milliseconds (cumulative count 999997)
100.000% <= 2.599 milliseconds (cumulative count 999999)
100.000% <= 2.623 milliseconds (cumulative count 1000000)
100.000% <= 2.623 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
98.247% <= 0.103 milliseconds (cumulative count 982469)
99.663% <= 0.207 milliseconds (cumulative count 996632)
99.889% <= 0.303 milliseconds (cumulative count 998894)
99.951% <= 0.407 milliseconds (cumulative count 999514)
99.988% <= 0.503 milliseconds (cumulative count 999884)
99.991% <= 0.607 milliseconds (cumulative count 999907)
99.995% <= 0.703 milliseconds (cumulative count 999946)
99.996% <= 0.807 milliseconds (cumulative count 999959)
99.997% <= 0.903 milliseconds (cumulative count 999971)
99.998% <= 1.007 milliseconds (cumulative count 999980)
100.000% <= 3.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 191350.94 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.060     0.016     0.055     0.087     0.127     2.623
====== GET ======                                                     
  1000000 requests completed in 5.13 seconds
  20 parallel clients
  1 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 2)
50.000% <= 0.055 milliseconds (cumulative count 598817)
75.000% <= 0.063 milliseconds (cumulative count 777031)
87.500% <= 0.071 milliseconds (cumulative count 892045)
93.750% <= 0.079 milliseconds (cumulative count 961078)
96.875% <= 0.087 milliseconds (cumulative count 978442)
98.438% <= 0.095 milliseconds (cumulative count 986643)
99.219% <= 0.111 milliseconds (cumulative count 992332)
99.609% <= 0.175 milliseconds (cumulative count 996251)
99.805% <= 0.239 milliseconds (cumulative count 998141)
99.902% <= 0.335 milliseconds (cumulative count 999028)
99.951% <= 0.431 milliseconds (cumulative count 999544)
99.976% <= 0.503 milliseconds (cumulative count 999765)
99.988% <= 0.687 milliseconds (cumulative count 999878)
99.994% <= 0.911 milliseconds (cumulative count 999945)
99.997% <= 1.487 milliseconds (cumulative count 999985)
99.999% <= 1.575 milliseconds (cumulative count 999993)
100.000% <= 1.583 milliseconds (cumulative count 1000000)
100.000% <= 1.583 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
99.049% <= 0.103 milliseconds (cumulative count 990495)
99.732% <= 0.207 milliseconds (cumulative count 997319)
99.882% <= 0.303 milliseconds (cumulative count 998823)
99.938% <= 0.407 milliseconds (cumulative count 999376)
99.977% <= 0.503 milliseconds (cumulative count 999765)
99.983% <= 0.607 milliseconds (cumulative count 999827)
99.989% <= 0.703 milliseconds (cumulative count 999886)
99.993% <= 0.807 milliseconds (cumulative count 999927)
99.994% <= 0.903 milliseconds (cumulative count 999937)
99.995% <= 1.007 milliseconds (cumulative count 999946)
99.996% <= 1.303 milliseconds (cumulative count 999957)
99.997% <= 1.407 milliseconds (cumulative count 999967)
99.999% <= 1.503 milliseconds (cumulative count 999989)
100.000% <= 1.607 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 194931.77 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.059     0.016     0.055     0.079     0.103     1.583
```

并发请求的Client 20个

发送总请求数为100000

value的大小为 5K字节

key的范围为10000内随机

```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -a redis -c 20 -n 1000000 -d 5k -r 100000 -t set,get
====== SET ======                                                     
  1000000 requests completed in 5.51 seconds
  20 parallel clients
  5 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 1)
50.000% <= 0.063 milliseconds (cumulative count 715664)
75.000% <= 0.071 milliseconds (cumulative count 762651)
87.500% <= 0.079 milliseconds (cumulative count 888364)
93.750% <= 0.087 milliseconds (cumulative count 940487)
96.875% <= 0.103 milliseconds (cumulative count 977615)
98.438% <= 0.111 milliseconds (cumulative count 984429)
99.219% <= 0.143 milliseconds (cumulative count 992489)
99.609% <= 0.199 milliseconds (cumulative count 996378)
99.805% <= 0.255 milliseconds (cumulative count 998147)
99.902% <= 0.311 milliseconds (cumulative count 999059)
99.951% <= 0.407 milliseconds (cumulative count 999534)
99.976% <= 0.447 milliseconds (cumulative count 999756)
99.988% <= 0.503 milliseconds (cumulative count 999883)
99.994% <= 0.567 milliseconds (cumulative count 999950)
99.997% <= 0.623 milliseconds (cumulative count 999970)
99.998% <= 0.647 milliseconds (cumulative count 999991)
99.999% <= 0.655 milliseconds (cumulative count 999997)
100.000% <= 0.663 milliseconds (cumulative count 999999)
100.000% <= 0.695 milliseconds (cumulative count 1000000)
100.000% <= 0.695 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
97.761% <= 0.103 milliseconds (cumulative count 977615)
99.672% <= 0.207 milliseconds (cumulative count 996715)
99.898% <= 0.303 milliseconds (cumulative count 998983)
99.953% <= 0.407 milliseconds (cumulative count 999534)
99.988% <= 0.503 milliseconds (cumulative count 999883)
99.996% <= 0.607 milliseconds (cumulative count 999965)
100.000% <= 0.703 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 181422.34 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.063     0.016     0.063     0.095     0.135     0.695
====== GET ======                                                     
  1000000 requests completed in 5.33 seconds
  20 parallel clients
  5 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.023 milliseconds (cumulative count 1)
50.000% <= 0.055 milliseconds (cumulative count 565972)
75.000% <= 0.071 milliseconds (cumulative count 837061)
87.500% <= 0.079 milliseconds (cumulative count 923520)
93.750% <= 0.087 milliseconds (cumulative count 954726)
96.875% <= 0.095 milliseconds (cumulative count 972474)
98.438% <= 0.111 milliseconds (cumulative count 988517)
99.219% <= 0.127 milliseconds (cumulative count 992281)
99.609% <= 0.191 milliseconds (cumulative count 996364)
99.805% <= 0.255 milliseconds (cumulative count 998202)
99.902% <= 0.319 milliseconds (cumulative count 999031)
99.951% <= 0.439 milliseconds (cumulative count 999536)
99.976% <= 0.591 milliseconds (cumulative count 999764)
99.988% <= 0.703 milliseconds (cumulative count 999881)
99.994% <= 0.871 milliseconds (cumulative count 999939)
99.997% <= 1.167 milliseconds (cumulative count 999970)
99.998% <= 1.463 milliseconds (cumulative count 999985)
99.999% <= 1.615 milliseconds (cumulative count 999997)
100.000% <= 1.623 milliseconds (cumulative count 999999)
100.000% <= 1.631 milliseconds (cumulative count 1000000)
100.000% <= 1.631 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
98.320% <= 0.103 milliseconds (cumulative count 983196)
99.692% <= 0.207 milliseconds (cumulative count 996923)
99.889% <= 0.303 milliseconds (cumulative count 998888)
99.938% <= 0.407 milliseconds (cumulative count 999384)
99.965% <= 0.503 milliseconds (cumulative count 999654)
99.979% <= 0.607 milliseconds (cumulative count 999786)
99.988% <= 0.703 milliseconds (cumulative count 999881)
99.991% <= 0.807 milliseconds (cumulative count 999907)
99.994% <= 0.903 milliseconds (cumulative count 999942)
99.997% <= 1.007 milliseconds (cumulative count 999969)
99.998% <= 1.207 milliseconds (cumulative count 999980)
99.999% <= 1.503 milliseconds (cumulative count 999990)
100.000% <= 1.703 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 187476.56 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.061     0.016     0.055     0.087     0.119     1.631

```

## 第二题

1. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

   

写入1万数据

```shell
127.0.0.1:6379> info Memory
# Memory
used_memory:893464
used_memory_human:872.52K
used_memory_rss:9973760
used_memory_rss_human:9.51M
used_memory_peak:56918608
used_memory_peak_human:54.28M
used_memory_peak_perc:1.57%
used_memory_overhead:851056
used_memory_startup:810056
used_memory_dataset:42408
used_memory_dataset_perc:50.84%
allocator_allocated:973736
allocator_active:1626112
allocator_resident:5038080
total_system_memory:3954188288
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.67
allocator_frag_bytes:652376
allocator_rss_ratio:3.10
allocator_rss_bytes:3411968
rss_overhead_ratio:1.98
rss_overhead_bytes:4935680
mem_fragmentation_ratio:11.70
mem_fragmentation_bytes:9121312
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:41000
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

127.0.0.1:6379> info Memory
# Memory
used_memory:1912496
used_memory_human:1.82M
used_memory_rss:10702848
used_memory_rss_human:10.21M
used_memory_peak:56918608
used_memory_peak_human:54.28M
used_memory_peak_perc:3.36%
used_memory_overhead:1382128
used_memory_startup:810056
used_memory_dataset:530368
used_memory_dataset_perc:48.11%
allocator_allocated:1935936
allocator_active:2318336
allocator_resident:5775360
total_system_memory:3954188288
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.20
allocator_frag_bytes:382400
allocator_rss_ratio:2.49
allocator_rss_bytes:3457024
rss_overhead_ratio:1.85
rss_overhead_bytes:4927488
mem_fragmentation_ratio:5.72
mem_fragmentation_bytes:8831368
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:41000
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入10万数据

```shell
127.0.0.1:6379> info memory
# Memory
used_memory:893464
used_memory_human:872.52K
used_memory_rss:10461184
used_memory_rss_human:9.98M
used_memory_peak:56918608
used_memory_peak_human:54.28M
used_memory_peak_perc:1.57%
used_memory_overhead:851048
used_memory_startup:810056
used_memory_dataset:42416
used_memory_dataset_perc:50.85%
allocator_allocated:940296
allocator_active:2109440
allocator_resident:5521408
total_system_memory:3954188288
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.24
allocator_frag_bytes:1169144
allocator_rss_ratio:2.62
allocator_rss_bytes:3411968
rss_overhead_ratio:1.89
rss_overhead_bytes:4939776
mem_fragmentation_ratio:12.27
mem_fragmentation_bytes:9608744
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:40992
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

127.0.0.1:6379> info memory
# Memory
used_memory:11532400
used_memory_human:11.00M
used_memory_rss:20205568
used_memory_rss_human:19.27M
used_memory_peak:56918608
used_memory_peak_human:54.28M
used_memory_peak_perc:20.26%
used_memory_overhead:5899632
used_memory_startup:810056
used_memory_dataset:5632768
used_memory_dataset_perc:52.53%
allocator_allocated:11552944
allocator_active:11808768
allocator_resident:15265792
total_system_memory:3954188288
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.02
allocator_frag_bytes:255824
allocator_rss_ratio:1.29
allocator_rss_bytes:3457024
rss_overhead_ratio:1.32
rss_overhead_bytes:4939776
mem_fragmentation_ratio:1.76
mem_fragmentation_bytes:8714184
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:41000
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50万数据

```shell
127.0.0.1:6379> info memory
# Memory
used_memory:893640
used_memory_human:872.70K
used_memory_rss:14585856
used_memory_rss_human:13.91M
used_memory_peak:56918608
used_memory_peak_human:54.28M
used_memory_peak_perc:1.57%
used_memory_overhead:851048
used_memory_startup:810056
used_memory_dataset:42592
used_memory_dataset_perc:50.96%
allocator_allocated:927352
allocator_active:6144000
allocator_resident:9601024
total_system_memory:3954188288
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:6.63
allocator_frag_bytes:5216648
allocator_rss_ratio:1.56
allocator_rss_bytes:3457024
rss_overhead_ratio:1.52
rss_overhead_bytes:4984832
mem_fragmentation_ratio:17.11
mem_fragmentation_bytes:13733240
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:40992
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

127.0.0.1:6379> info memory
# Memory
used_memory:53374944
used_memory_human:50.90M
used_memory_rss:62038016
used_memory_rss_human:59.16M
used_memory_peak:56918608
used_memory_peak_human:54.28M
used_memory_peak_perc:93.77%
used_memory_overhead:25045360
used_memory_startup:810056
used_memory_dataset:28329584
used_memory_dataset_perc:53.89%
allocator_allocated:53398368
allocator_active:53661696
allocator_resident:57073664
total_system_memory:3954188288
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.00
allocator_frag_bytes:263328
allocator_rss_ratio:1.06
allocator_rss_bytes:3411968
rss_overhead_ratio:1.09
rss_overhead_bytes:4964352
mem_fragmentation_ratio:1.16
mem_fragmentation_bytes:8704088
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:41000
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

查询key占用命令

```shell
./redis-cli -p 6379 -a redis --bigkeys

Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.

# Scanning the entire keyspace to find biggest keys as well as
# average sizes per key type.  You can use -i 0.1 to sleep 0.1 sec
# per 100 SCAN commands (not usually needed).

[00.00%] Biggest string found so far '"redis_key_167023"' with 38 bytes
[00.01%] Biggest string found so far '"redis_key_39546"' with 39 bytes

-------- summary -------

Sampled 500000 keys in the keyspace!
Total key length in bytes is 7888890 (avg len 15.78)

Biggest string found '"redis_key_39546"' has 39 bytes

0 lists with 0 items (00.00% of keys, avg size 0.00)
0 hashs with 0 fields (00.00% of keys, avg size 0.00)
500000 strings with 12547288 bytes (100.00% of keys, avg size 25.09)
0 streams with 0 entries (00.00% of keys, avg size 0.00)
0 sets with 0 members (00.00% of keys, avg size 0.00)
0 zsets with 0 members (00.00% of keys, avg size 0.00)
```

500000 strings with 12547288 bytes (100.00% of keys, avg size 25.09)

50万的key 占用 12547288 bytes

50万的key平均大小为 25.09 bytes

50万的value 占用 39934016 bytes































