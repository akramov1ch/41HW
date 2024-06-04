package handlers

import (
    "net/http"
    "sync"
    "github.com/gin-gonic/gin"
)

var (
    requestCount = make(map[string]int)
    mu           sync.Mutex
)

func RequestHandler(c *gin.Context) {
    clientIP := c.ClientIP()

    mu.Lock()
    requestCount[clientIP]++
    count := requestCount[clientIP]
    mu.Unlock()

    if count >= 3 {
        c.JSON(http.StatusForbidden, gin.H{"message": "Maximum attempts reached"})
        return
    }
    desiredWord := "hello"
    receivedWord := c.Query("word")

    if receivedWord == desiredWord {
        c.JSON(http.StatusOK, gin.H{"message": "Correct word received!"})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Try again"})
    }
}
