package main

import (
	"net/http"
	"os"
    "fmt"
    "io"
)

type logWriter struct {}

func main() {
    resp, err := http.Get("http://google.com")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    lw := logWriter{}
    io.Copy(lw, resp.Body)
}

// log writer now implements the Writer interface,
// because it implements a Write function
// that takes in a byte slice and returns an int and an error.
func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Println("Just wrote this many bytes:", len(bs))
    return len(bs), nil
}
