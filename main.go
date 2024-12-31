package main

import (
    "fmt"
    "runtime"
    "crypto/rand"
    "net"
)

func main() {
    fmt.Println("hello world ",runtime.Version())

    b := make([]byte, 8)
    reader := rand.Reader
    n, rerr := reader.Read(b)
    if rerr != nil {
        fmt.Println("rand error ", rerr)
    }
    fmt.Println("rand bytes num: ", n," b:", b)
    conn, err := net.Dial("tcp", "142.251.40.145:http")

    if err != nil {
        fmt.Println("dial error ", err)
    }

    conn.Close()
}
