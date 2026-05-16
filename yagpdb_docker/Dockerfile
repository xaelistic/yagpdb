# Everdream fork of YAGPDB
# Build: docker compose -f yagpdb_docker/docker-compose.yml build
# Run:   docker compose -f yagpdb_docker/docker-compose.yml up -d

FROM docker.io/golang:1.26.0 AS builder

WORKDIR /appbuild/yagpdb
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /appbuild/yagpdb/cmd/yagpdb
RUN GOEXPERIMENT=jsonv2 CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-X github.com/xaelistic/yagpdb/v2/common.VERSION=$(git describe --tags)-everdream"

FROM docker.io/alpine:latest

LABEL org.opencontainers.image.source="https://github.com/xaelistic/yagpdb"
LABEL org.opencontainers.image.description="Everdream Discord Bot (YAGPDB fork)"

WORKDIR /app
VOLUME ["/app/soundboard", "/app/cert"]
EXPOSE 80 443

RUN apk --no-cache add ca-certificates ffmpeg tzdata

COPY --from=builder /appbuild/yagpdb/cmd/yagpdb/yagpdb yagpdb

ENTRYPOINT ["/app/yagpdb"]
CMD ["-all", "-pa", "-exthttps=false", "-https=true"]
