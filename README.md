# Edge-Caching Reverse Proxy
A reverse proxy written in Go that caches backend responses. It works by using a cache 
system implemented by using a linked list to control proper caching logic and a GO
map to track the clients URL's and related responses.

The map witholds the clients request url and the pointer to the location of the related 
response within the linked list.

## Status
Work in progress — The proxy itself is finished and the cache has been implmented.

## Why
I wanted to challenge myself by stepping into unknown waters like
creating a proxy and learning GO. Gives me basic insite to caching 
strategies.

## Planned features
- Implement the proper system for allowing different request types and responses, at the moment the proxy only follows 
a hardcoded path from the server.
- Implementing mutexing on the proxy to prevent concurrency errors.
- Implementing TTL-based cache expiry to follow proper caching logic.
- Creating tests and gathering results.

## Running locally

### Prerequisites
- Go 1.2x or later installed

### Project structure

This repo includes three components:
- **`proxy`** — the caching reverse proxy (LRU cache + handler)
- **`server`** — a mock backend server for testing (responds "hello world" after an artificial 5-second delay, to make the caching benefit obvious)
- **`client`** — a simple test client that sends a request and prints the response/status

### Configuration

Currently hardcoded in source rather than configurable via flags:
- **Backend URL**: set in the proxy's `handler()` as `http://localhost:8080`
- **Proxy listen port**: `9090`
- **Cache capacity**: set via the `NewLRU(cap)` call in the proxy's `main.go`

### Running the demo

1. Start the mock backend server:
```bash
   go run server/main.go
```
2. In a separate terminal, start the proxy:
```bash
   go run proxy/main.go
```
3. In a third terminal, run the client:
```bash
   go run client/main.go
```

The first request will take ~5 seconds (hitting the backend directly). Run the client again for the same path and it should return near-instantly, served from the LRU cache.

## Acknowledgments
- AI assistance (Claude) was used for debugging and documentation support during development.
- Built with [Go](https://go.dev/) and the standard library [`container/list`](https://pkg.go.dev/ container/list) package for the LRU cache's linked list implementation.
- proxy vs reverse proxy: https://youtu.be/4NB0NDtOwIQ?si=v3og3r69dGOsbb0V
- How to write a reverse proxy with GO: https://youtu.be/vlPCAEUCCa0?si=oHocVM4wNlZoiy7D
- How to write an LRU cache in go: https://youtu.be/DvmMYD8oaZw?si=0UYhiJKiFXxnF53x

