From golang:1.20.6
WORKDIR /app
COPY . .
RUN go run main.go -o server
CMD ["/app/server"]