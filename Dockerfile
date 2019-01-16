# start from golang image based on alpine-3.8
FROM golang:1.11-alpine AS dev-build
# add our cgo dependencies
RUN apk add --no-cache ca-certificates cmake make g++ openssl-dev git curl pkgconfig
# clone seabolt source code
RUN git clone -b 1.7 https://github.com/neo4j-drivers/seabolt.git /seabolt
# invoke cmake build and install artifacts - default location is /usr/local
WORKDIR /seabolt/build
# CMAKE_INSTALL_LIBDIR=lib is a hack where we override default lib64 to lib to workaround a defect
# in our generated pkg-config file
RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_LIBDIR=lib .. && cmake --build . --target install

ENV GO111MODULE on
WORKDIR /go/src/neo4j-go-file-system
ADD . .
# install dependencies
RUN go get -v ./...
# build statically linked executable
RUN go build -o app -tags seabolt_static main.go

# base the image to neo4j
FROM neo4j:latest
# copy the statically linked executable
COPY --from=dev-build /go/src/neo4j-go-file-system/app /neo4j-go-file-system/app