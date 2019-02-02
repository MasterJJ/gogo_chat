package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
)

const DEF_MAX_LEN = 1024
const DEF_DEFAULT_HOST = ""
const DEF_DEFAULT_PORT = ""


var  writeStr, readStr = make([]byte, DEF_MAX_LEN), make([]byte, DEF_MAX_LEN)

func main() {
        var (
                host = DEF_DEFAULT_HOST
                port = DEF_DEFAULT_PORT
                remote = host + ":" + port
                reader = bufio.NewReader(os.Stdin)
        )


}
