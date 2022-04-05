FROM golang:latest AS builder
RUN git clone --progress --verbose --depth=1 https://github.com/Bpazy/behappy /behappy
WORKDIR /behappy
RUN go env && CGO_ENABLED=0 go build ./cmd/behappy

FROM alpine:latest AS production
COPY --from=builder /behappy/behappy /behappy/behappy
WORKDIR /behappy
ENTRYPOINT ["./behappy"]
