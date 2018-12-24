ARG GO_VERSION

# builder image
FROM golang:$GO_VERSION-alpine as builder

ENV GO111MODULE=on
RUN apk add -U --no-cache curl build-base git
WORKDIR /go/src/github.com/gofunct/hack/tools
COPY . .

RUN GO111MODULE=on go install ./cmd/tools


# test image
FROM golang:$GO_VERSION-alpine

ENV DEP_RELEASE_TAG v0.5.0
RUN apk add -U --no-cache curl build-base git
COPY --from=builder /go/bin/tools /go/bin/

WORKDIR /go/src/myapp
