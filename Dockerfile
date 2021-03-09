FROM golang:1.13.5-alpine
RUN apk add --no-cache git make
#     go get "github.com/lib/pq" \
#            "github.com/caarlos0/env" \
#            "github.com/prometheus/client_golang/prometheus" \
#            "github.com/prometheus/client_golang/prometheus/promhttp"
WORKDIR /go/src/app/
COPY . .
RUN make

FROM alpine:3.10.3
WORKDIR /app
COPY --from=0 /go/bin/apiserver .
ENTRYPOINT ["./apiserver"]