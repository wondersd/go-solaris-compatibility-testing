# Solaris 11.2 Go Compat Testing #

## Building ##

```
podman build . -t hello-world && podman run -d --rm --name helloworld -it --entrypoint cat localhost/hello-world:latest && podman cp helloworld:/usr/local/bin/hello-world . && podman kill helloworld
```

## Tests ## 

The three things that my project needed and proved to 

* hello world/entrypoint
* cryto/rand
* net/Dial

## Solaris 11.2 ##

Solaris 11.2 Live CD

```
jack@solaris:~$ uname -a
SunOS solaris 5.11 11.2 i86pc i386 i86pc
```

## 1.21 - current (alpine, musl-libc) ##

https://github.com/golang/go/issues/61950

Builder Images Tested

`golang:1.24rc1-alpine3.20`
`golang:1.23-alpine3.20`
`golang:1.22-alpine3.21`
`golang:1.21-alpine3.20`
`golang:1.21-bullseye`

seg fault on start

```
jack@solaris:~$ /jack/hello-world 
Segmentation Fault (core dumped)
jack@solaris:~$ /jack/hello-world 
Segmentation Fault (core dumped)
jack@solaris:~$ /jack/hello-world 
Segmentation Fault (core dumped)
jack@solaris:~$ /jack/hello-world 
hello world  go1.20.14
```

## 1.16 compiled on debian (glibc) ##

Builder Images Tested

`golang:1.16-bullseye`
`golang:1.13.15-buster`

So far the latest compatible combination to cross compile on solaris

```
jack@solaris:~$ /jack/hello-world 
hello-world: Cannot find /lib/ld-musl-x86_64.so.1
Killed
jack@solaris:~$ /jack/hello-world 
hello world  go1.13.15
rand bytes num:  8  b: [211 233 99 30 63 210 128 11]
jack@solaris:~$ /jack/hello-world 
ld.so.1: hello-world: fatal: relocation error: file /jack/hello-world: symbol getrandom: referenced symbol not found
Killed
jack@solaris:~$ /jack/hello-world 
hello world  go1.17.13
ld.so.1: hello-world: fatal: relocation error: file /jack/hello-world: symbol getrandom: referenced symbol not found
Killed
jack@solaris:~$ /jack/hello-world 
hello world  go1.16.15
rand bytes num:  8  b: [95 39 216 112 142 167 47 192]
```

## 1.17 compiled on debian (glibc) ##

Builder Images Tested

`golang:1.17-bullseye`

```
jack@solaris:~$ /jack/hello-world 
hello world  go1.17.13
ld.so.1: hello-world: fatal: relocation error: file /jack/hello-world: symbol getrandom: referenced symbol not found
Killed
```

## 1.18 (alpine, musl-libc) ##

last version to have a working network stack

`golang:1.18-alpine3.17`
`golang:1.19-alpine3.18`
`golang:1.20-alpine3.18`

```
jack@solaris:~$ /jack/hello-world 
hello world  go1.20.14
dial error  dial tcp 142.251.40.145:80: socket: protocol wrong type for socket
jack@solaris:~$ /jack/hello-world 
hello world  go1.19.13
dial error  dial tcp 142.251.40.145:80: socket: protocol wrong type for socket
jack@solaris:~$ /jack/hello-world 
hello world  go1.19.13
dial error  dial tcp 142.251.40.145:80: socket: protocol wrong type for socket
jack@solaris:~$ /jack/hello-world 
hello world  go1.19.13
dial error  dial tcp 142.251.40.145:80: socket: protocol wrong type for socket
jack@solaris:~$ /jack/hello-world 
hello world  go1.18.10
```



