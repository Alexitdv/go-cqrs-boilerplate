FROM golang:1.15-alpine AS build

ARG GOOS=linux
ARG PORT=5000
ARG HTTP_PORT=8080
ARG APP_PKG_NAME=wb-cards

ENV CGO_ENABLED=0 \
    GOOS=$GOOS \
    GOARCH=amd64 \
    CGO_CPPFLAGS="-I/usr/include" \
    UID=0 GID=0 \
    CGO_CFLAGS="-I/usr/include" \
    CGO_LDFLAGS="-L/usr/lib -lpthread -lrt -lstdc++ -lm -lc -lgcc " \
    PKG_CONFIG_PATH="/usr/lib/pkgconfig" \
    GO111MODULE=on \
    APP_PKG_NAME=$APP_PKG_NAME \
    PORT=$PORT

RUN apk update && apk add --no-cache curl protoc musl-dev gcc git build-base make bash

WORKDIR /go/src/$APP_PKG_NAME
COPY . .

#RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.38.0
#RUN golangci-lint run
#RUN go test ./...

RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go \
              google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN /bin/bash ./scripts/gen.sh

ARG HASHCOMMIT
ARG VERSION

RUN go build -mod=mod -v -o /out/migration ./tools/migration/*.go
RUN go build -mod=mod -v \
    -o /out/auth \
    -ldflags "-extldflags '-static' -X 'main.serviceVersion=$VERSION' -X 'main.hashCommit=$HASHCOMMIT'" \
    ./cmd/auth/*.go


# Copy to Alpine image
FROM alpine:3.12
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /out/auth /app/auth
COPY --from=build /out/migration /app/migration
COPY ./configs/keys /app/configs/keys
COPY ./migrations /app/migrations
EXPOSE 80 $HTTP_PORT $PORT
CMD /app/migration -dir /app/migrations/sql up && /app/auth
