ARG GOLANG_VERSION=1.14
FROM golang:${GOLANG_VERSION}-alpine as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY version/ version/
COPY pkg/ pkg/

# Build
ARG GIT_COMMIT
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build -a -o manager -ldflags -X=github.com/configurator/multitenancy/version.CommitSHA=${GIT_COMMIT} \
    cmd/manager/main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/multitenancy \
    USER_UID=1001 \
    USER_NAME=multitenancy

# install operator binary
COPY --from=builder /workspace/manager ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
