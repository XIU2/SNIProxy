FROM golang:1.23-alpine AS builder
WORKDIR /workspace

COPY . .
RUN go get
RUN go build -o /workspace/sniproxy


FROM golang:1.23-alpine AS release

WORKDIR /sniproxy

RUN mkdir -p /sniproxy/conf

# Copy from builder
COPY --from=builder /workspace/sniproxy ./sniproxy

RUN chmod +x sniproxy
COPY config.yaml ./conf

EXPOSE 443

CMD ["./sniproxy","-c","./conf/config.yaml"]