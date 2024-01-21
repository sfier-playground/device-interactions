# Defining App builder image
FROM golang:1.21-alpine3.17 AS builder

ARG PAT
RUN apk update; \
    apk add --no-cache \
    git \
    make
RUN git config --global url."https://${PAT}:sifer169966@github.com/".insteadOf "https://github.com/"
RUN go env -w GOPRIVATE=github.com/sifer169966

# Define current working directory
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . ./

# Build App
RUN make build

# Defining App image
FROM alpine:3.17 as release

RUN apk add --no-cache --update ca-certificates

# Copy App binary to image
COPY --from=builder /app/device-interaction /app/cmd/

RUN chmod +x /app/cmd/device-interaction

ARG NONROOT_GROUP=nonroot-group
ARG NONROOT_USER=nonroot-user
ARG USER_ID=20000

RUN addgroup -S $NONROOT_GROUP && adduser -S -u $USER_ID $NONROOT_USER -G $NONROOT_GROUP

USER $NONROOT_USER:$NONROOT_GROUP

WORKDIR /app

CMD ["cmd/device-interaction", "serve-rest"]
