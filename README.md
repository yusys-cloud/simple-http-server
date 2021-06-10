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

