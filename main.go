package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    engin := gin.Default()
    
    DefineRoutes(engin)
    engin.Run(":3000")
}

func DefineRoutes(r gin.IRouter) {
    r.GET("/", func(ctx *gin.Context) {
        ctx.JSONP(http.StatusOK, gin.H{
            "message": "ok",
            "data": "Hello World",
        })
    })
}
