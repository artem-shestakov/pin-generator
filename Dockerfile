FROM golang:1.13.5-alpine
RUN apk add --no-cache git make gcc g++
WORKDIR /go/src/app/
COPY . .
RUN make -f ./scripts/Makefile

FROM alpine:3.10.3
WORKDIR /app
COPY ./configs/ /app/configs/
COPY --from=0 /go/src/app/api/swagger.yml ./api/
COPY --from=0 /go/src/app/apiserver .
EXPOSE 8080
ENTRYPOINT ["./apiserver"]