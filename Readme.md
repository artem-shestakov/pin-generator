# Pin code generator
Generate pin code, salt and SHA-1 hash.

## Build
Use make file to build application:
```
make -f ./scripts/Makefile
```

## Docs
Start server and open http://server_ip:8080/docs

To generate swagger.yaml use Makefile:
```
make swagger -f ./scripts/Makefile 
```

## Docker
Docker image [here](https://hub.docker.com/repository/docker/artemshestakov/pin-generator/general).