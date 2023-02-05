FROM golang:1.19-alpine AS build

WORKDIR /app
COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go install github.com/cosmtrek/air@v1.27.3

# usermodなどで手元のUIDが変わっている場合は.envに記述する
RUN chown -R ${UID:-1000}:${GID:-1000} ./

CMD ["air","-c",".air.toml"]