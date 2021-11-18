# Simple Http Server

a simple http server.

## Quick Start

``` 
./simple-http-server [-d DIRECTORY] [-p PORT]
```

## Examples
Serving static files 
``` 
http://localhost:3000/assets
```

Serve from your current directory:

``` 
./simple-http-server
```

Serve from another directory:

``` 
./simple-http-server -d /another/directory
```

Serve from port 1234:

```
./simple-http-server -p 1234

```

## Benchmarks
| Benchmark name                 |       (1) |
| ------------------------------ | --------- | 
| Benchmark_mac_ping             | 1012/sec  |

summary = 480903 in 00:05:10 = 1552.7/s Avg:     2 Min:     2 Max:    37 Err:     0 (0.00%)
summary +  47700 in 00:00:30 = 1590.4/s Avg:     2 Min:     2 Max:    11 Err:     0 (0.00%) Active: 4 Started: 4 Finished: 0
summary = 528603 in 00:05:40 = 1556.0/s Avg:     2 Min:     2 Max:    37 Err:     0 (0.00%)
summary +  48100 in 00:00:30 = 1600.6/s Avg:     2 Min:     1 Max:    23 Err:     0 (0.00%) Active: 4 Started: 4 Finished: 0
summary = 576703 in 00:06:10 = 1559.6/s Avg:     2 Min:     1 Max:    37 Err:     0 (0.00%)
summary +  49600 in 00:00:30 = 1656.4/s Avg:     2 Min:     2 Max:    13 Err:     0 (0.00%) Active: 4 Started: 4 Finished: 0
summary = 626303 in 00:06:40 = 1566.9/s Avg:     2 Min:     1 Max:    37 Err:     0 (0.00%)
summary +  48000 in 00:00:30 = 1600.0/s Avg:     2 Min:     2 Max:    30 Err:     0 (0.00%) Active: 4 Started: 4 Finished: 0
summary = 674303 in 00:07:10 = 1569.2/s Avg:     2 Min:     1 Max:    37 Err:     0 (0.00%)
summary +  45200 in 00:00:30 = 1502.4/s Avg:     2 Min:     2 Max:    14 Err:     0 (0.00%) Active: 4 Started: 4 Finished: 0

