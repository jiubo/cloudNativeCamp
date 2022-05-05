FROM golang:1.18 as builder

WORKDIR /workspace

COPY /cmd/httpServer/main.go .

RUN CGO_ENABLED=0 go build -o server main.go

FROM scratch

WORKDIR /

COPY --from=builder /workspace/server /server

ENTRYPOINT ["/server"]