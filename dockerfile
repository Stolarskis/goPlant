FROM golang:latest
WORKDIR /home
COPY . ./ 
RUN go build ./cmd/api/main.go 
ENTRYPOINT ["./main"]