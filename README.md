# Memgo - An in-memory store

Memgo is a simple in-memory store written with Go and gRPC.

## Running in Docker

```bash
```

## Runtime Options

Runtime options are set as environment variables with default values as fallbacks.

### Port

```
MEMGO_PORT=8000
```

### Username

```
MEMGO_USER=memgo
```

### Password

```
MEMGO_PASSWORD=memgo
```

### Max Databse Size

This is a number, followed by `M` for Megabytes or `G` for Gigabytes.

```
MEMGO_MAX_SIZE=128M
```

## Best Practices

* **Set an expiry time** - this will ensure that data doesn't accumulate unnecessarily.
* **Don't expose port** - if running without TLS enabled, keep the port only exposed within your application network.
