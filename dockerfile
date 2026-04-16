#Base image 
FROM golang:1.26.1 AS builder

#working dir for docker
WORKDIR /app

#copy mod files first
COPY go.mod go.sum ./

#Install dependencies
RUN go mod download

#copy the whole project in /app dir
COPY . .

#Build app
RUN go build -o main ./cmd

#introduce the lighweight environment
FROM debian:bookworm-slim

#working dir that needs to be run just created above
WORKDIR /app

# 🔥
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

#copy the final compiled binary only not build tools and other overheads
COPY --from=builder /app/main .

#expose app port
EXPOSE 8080

#RUN the compiled binary i.e main file
CMD ["./main"]
