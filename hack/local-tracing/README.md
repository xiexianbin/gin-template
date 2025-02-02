# Local tracing setup

ref https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.101.0/examples/demo

## start

```
docker compose up -d
```

exposes the following backends:

- Jaeger at http://0.0.0.0:16686
- Zipkin at http://0.0.0.0:9411
- Prometheus at http://0.0.0.0:9090
