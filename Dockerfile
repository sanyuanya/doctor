FROM golang:1.24.5 AS builder

WORKDIR /app

COPY . .


RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -gcflags=all="-N -l" -o /task

FROM debian:stable-slim

COPY --from=builder /task /task

ENV TZ=Asia/Shanghai \
  DEBIAN_FRONTEND=noninteractive

ENV ENVIRONMENT=production

RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
  && echo ${TZ} > /etc/timezone \
  && dpkg-reconfigure --frontend noninteractive tzdata \
  && rm -rf /var/lib/apt/lists/*

ENTRYPOINT [ "/task" ]