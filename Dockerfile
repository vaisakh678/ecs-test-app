FROM golang:1.22.5-alpine


WORKDIR /app
RUN apk add --no-cache git
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o server .
EXPOSE 3000

CMD ["./server"]
