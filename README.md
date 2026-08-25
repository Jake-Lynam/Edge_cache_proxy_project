# Edge-Caching Reverse Proxy

A reverse proxy written in Go that caches backend responses to reduce 
latency on repeated requests.

## Status
Work in progress — currently built the origin server and working on
core proxy logic.

## Why
I wanted to challenge myself by stepping into unknown waters like
creating a proxy and learning GO. Gives me basic insite to caching 
strategies and concurrency handling

## Planned features
- Reverse proxying to a backend origin server
- TTL-based cache expiry
- LRU eviction policy
- Concurrent request handling / stampede protection
- Metrics endpoint (hit/miss ratio, latency)
- Automated tests
- Benchmark results

## Running locally
(fill in once there's something runnable end-to-end)