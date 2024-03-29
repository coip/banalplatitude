FROM golang:alpine as builder
WORKDIR /usr/src/bp
ADD main.go .
ADD go.mod .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /listener .
RUN echo "nobody:x:65534:65534:nobody:/:/sbin/nologin" >> /tmp/nobody

# final stage
FROM scratch
COPY --from=builder /tmp/nobody /etc/passwd
COPY --from=builder /listener /listener
USER nobody
ENTRYPOINT ["/listener"]
