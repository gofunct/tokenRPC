ARG alpine_version=3.8
FROM golang:1.11-alpine$alpine_version AS build

RUN apk add --update --no-cache git
WORKDIR /app
ENV GOPATH=/app

# Copy all of the staged files (protos plus go source)
COPY . /app/

# Download the go dependencies.
RUN go get ./...

WORKDIR /app

# Build the gateway
RUN go build -o tokenRPC .

FROM alpine:$alpine_version
WORKDIR /app
COPY --from=build /app/tokenRPC /app/
COPY --from=build /app/tokenRPC.yaml /app/
COPY --from=build /app/contracts/token_service/token_service.swagger.json /app/

EXPOSE 8080
ENTRYPOINT ["/app/tokenRPC", "gateway"]
