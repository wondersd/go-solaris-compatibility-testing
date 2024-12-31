# build
# need debian based builder image so that its
# linked against glibc and not musl-libc as solaris 11.2
# doesnt have that by default
FROM golang:1.16-bullseye AS builder

WORKDIR /work

COPY go.mod ./
RUN go mod download

COPY main.go .
RUN GOOS=solaris GOARCH=amd64 go build -o bin/hello-world

FROM alpine:3.20

COPY --from=builder /work/bin/hello-world /usr/local/bin/
ENTRYPOINT ["hello-world"]