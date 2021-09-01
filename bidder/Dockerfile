
FROM golang:1.16 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o bidder ./main.go

FROM ubuntu:20.04
RUN apt update; \
        export DEBIAN_FRONTEND=noninteractive; \
        apt-get install -y tzdata; \
        dpkg-reconfigure --frontend noninteractive tzdata; \
        ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime; \
        rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/bidder  .
COPY --from=builder /app/.env  .

EXPOSE 7000

CMD ["./bidder"]