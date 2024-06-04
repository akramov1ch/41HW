package main

import (
    "server/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}
