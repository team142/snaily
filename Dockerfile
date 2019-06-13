#
# Multi-stage build Go app
#

# Build Golang app in Go container
FROM golang:1.12.4-alpine3.9 as builder
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates git

FROM builder as builder2
WORKDIR /src
COPY go.mod ./
RUN go mod download

FROM builder2 as builder3
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final app
FROM scratch as final
COPY --from=builder3 /user/group /user/passwd /etc/
COPY --from=builder3 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder3 /src/app /app
USER nobody:nobody
ENTRYPOINT ["/app", "-addrToProxy=http://snaily-static:80", "-pgport=5432", "-pghost=spg", "-nsqd=nsq_nsqd:4150"]