FROM golang:1.15.6-alpine AS build

WORKDIR /src/
COPY primes.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]