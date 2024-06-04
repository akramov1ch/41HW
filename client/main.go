package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func main() {
    for i := 0; i < 3; i++ {
        resp, err := http.Get("http://localhost:8080/request?word=wrongword")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
            return
        }

        fmt.Println("Response:", string(body))

        if resp.StatusCode == http.StatusOK {
            fmt.Println("Received the desired response, stopping...")
            return
        }

        time.Sleep(2 * time.Second)
    }
}
