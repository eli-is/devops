# B"H
FROM golang:1.21-alpine as build
# Set the Current Working Directory inside the container
WORKDIR /app/go-http-server

# download Go modules and dependencies
COPY go.mod .
RUN go mod download

# Build the Go app
COPY . .
RUN go build -o ./out/go-http-server .

# Run the app in a small container as non root user
FROM alpine:3.18

WORKDIR /
COPY --from=build /app/go-http-server/out/go-http-server /go-http-server

RUN adduser -D app
USER app

# Add in case of debugging 
ARG GITHUB_SHA=""
ARG GITHUB_REF=""

ENV GITHUB_SHA=${GITHUB_SHA}
ENV GITHUB_REF=${GITHUB_REF}

# Run the binary program produced by `go install`
CMD ["/go-http-server"]