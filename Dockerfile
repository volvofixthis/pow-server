FROM registry.lux.cloud.gc.onl/docker-io/golang:1.23.1 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY internal/ internal/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o tcpclient cmd/tcpclient/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/server .
COPY --from=builder /workspace/tcpclient .
USER 65532:65532

ENTRYPOINT ["/server"]
