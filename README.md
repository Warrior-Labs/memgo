# Memgo - An in-memory store

Memgo is a simple in-memory store written with Go and gRPC.

## Running in Docker

```bash
docker run -d --name memgo
```

## Using TLS

To enable TLS, you must specify a Docker volume which contains a Certificate called `certificate.pem` and a Key called `key.pem`. These should be mounted to the `/x509` directory in the container:

```bash
docker run -d --name memgo -v /absolute/path/to/certs:/x509
```

To confirm TLS is running, you can check the logs

```bash
docker logs memgo
```

This should output the following:

```bash
TLS Enabled
Starting Memgo Server on port 8000
```

## Runtime Options

Runtime options are set as environment variables with default values as fallbacks.

### Max Database Size

This is a number, followed by `M` for Megabytes or `G` for Gigabytes. The defaut is shown below:

```
MEMGO_MAX_SIZE=128M
```

## Best Practices

* **Set an expiry time** - this will ensure that data doesn't accumulate unnecessarily.
* **Don't expose port** - if running without TLS enabled, keep the port only exposed within your application network.
