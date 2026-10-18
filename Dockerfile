FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cahce git

ENV GOPRIVATE=github.com/SamuraiJeka/SnapLink-proto

COPY go.mod go.sum ./

RUN --mount=type=secret,id=netrc,target=/root/.netrc,required=true \
    go mod download

COPY . .

RUN go build -o /out/service ./cmd/main.go


FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && adduser -D -H appuser

WORKDIR /app

COPY --from=builder /out/service ./service

EXPOSE 8080

ENTRYPOINT [ "./service" ]