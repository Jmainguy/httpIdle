# HTTP Idle Server

This project is a simple Go HTTP server that demonstrates idle connection handling. The server responds with a "hello" message and then pauses for a configured idle time.

## Features

- Configurable idle wait time (time the server sleeps after responding).
- Configurable server address and keep-alive behavior.
- Can be run as an RPM package or via a Docker container.

## Configuration

The server uses environment variables to configure the behavior:

- `HTTPIDLE_IDLEWAIT`: Time in seconds to wait after sending a response (default: 10 seconds).
- `HTTPIDLE_KEEPALIVE`: Enables/disables HTTP keep-alive (default: true).
- `HTTPIDLE_ADDR`: Address and port to listen on (default: ":8080").

## Installation

### RPM Installation

You can download the latest RPM package from the [GitHub releases page](https://github.com/jmainguy/httpIdle/releases).

### Docker Installation
A pre-built Docker image is available at hub.soh.re/httpidle:v1.0.0. You can pull and run the container as follows:

```bash
docker pull hub.soh.re/httpIdle:v1.0.0
docker run -d -p 8080:8080 --name httpIdle hub.soh.re/httpIdle:v1.0.0
```
The server will start on the configured address (default: :8080), and you can test it by sending a request:

```bash
curl http://localhost:8080
```

### Environment Variables Example
```bash
export HTTPIDLE_IDLEWAIT=5
export HTTPIDLE_KEEPALIVE=false
export HTTPIDLE_ADDR=":9090"
httpIdle
```
