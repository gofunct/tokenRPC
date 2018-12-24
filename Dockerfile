ARG alpine=3.8
ARG go=1.11.0
ARG grpc
ARG grpc_java
FROM golang:$go-alpine$alpine AS build
# TIL docker arg variables need to be redefined in each build stage
ARG grpc
ARG grpc_java
RUN set -ex && apk --update --no-cache add \
    bash \
    make \
    cmake \
    autoconf \
    automake \
    curl \
    tar \
    libtool \
    g++ \
    git \
    openjdk8-jre \
    libstdc++ \
    ca-certificates

WORKDIR /build
COPY install-protobuf.sh /build
RUN chmod +x /build/install-protobuf.sh
RUN /build/install-protobuf.sh ${grpc} ${grpc_java}
RUN git clone https://github.com/googleapis/googleapis

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/vX.X.X/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN curl -sSL https://github.com/uber/prototool/releases/download/v1.3.0/prototool-$(uname -s)-$(uname -m) \
    -o /usr/local/bin/prototool && \
    chmod +x /usr/local/bin/prototool

# Add grpc-web support

RUN curl -sSL https://github.com/grpc/grpc-web/releases/download/1.0.3/protoc-gen-grpc-web-1.0.3-linux-x86_64 \
    -o /build/grpc_web_plugin && \
    chmod +x /build/grpc_web_plugin

RUN go get -u \
        		github.com/getamis/sol2proto \
        		github.com/getamis/grpc-contract \
        		github.com/ethereum/go-ethereum/cmd/abigen \
        		google.golang.org/grpc \
        		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
        		github.com/golang/protobuf/protoc-gen-go \
        		github.com/ckaznocha/protoc-gen-lint \
        		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
        		moul.io/protoc-gen-gotemplate \
        		github.com/gogo/protobuf/...

FROM alpine:$alpine AS protoc-all

RUN set -ex && apk --update --no-cache add \
    bash \
    libstdc++ \
    libc6-compat \
    ca-certificates

COPY --from=build /build/grpc/bins/opt/grpc_* /usr/local/bin/
COPY --from=build /build/grpc/bins/opt/protobuf/protoc /usr/local/bin/
COPY --from=build /build/grpc/libs/opt/ /usr/local/lib/
COPY --from=build /build/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java /usr/local/bin/
COPY --from=build /build/googleapis/google/ /usr/local/include/google
COPY --from=build /usr/local/include/google/ /usr/local/include/google
COPY --from=build /usr/local/bin/prototool /usr/local/bin/prototool
COPY --from=build /go/bin/* /usr/local/bin/
COPY --from=build /go/src/* /usr/local/include/
COPY --from=build /build/grpc_web_plugin /usr/local/bin/grpc_web_plugin
COPY --from=build /go/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options/ /usr/local/include/protoc-gen-swagger/options/

ADD entrypoint.sh /usr/local/bin
RUN chmod +x /usr/local/bin/entrypoint.sh

WORKDIR /hack
ENTRYPOINT [ "entrypoint.sh" ]

# protoc
FROM protoc-all AS protoc
ENTRYPOINT [ "protoc", "-I/usr/local/include" ]

# prototool
FROM protoc-all AS prototool
ENTRYPOINT [ "prototool" ]

# grpc-cli
FROM protoc-all as grpc-cli

# gen-grpc-gateway
FROM protoc-all AS gen-templates

COPY templates /templates
COPY generate_gateway.sh /usr/local/bin
RUN chmod +x /usr/local/bin/generate_gateway.sh

WORKDIR /hack
ENTRYPOINT [ "generate_gateway.sh" ]
