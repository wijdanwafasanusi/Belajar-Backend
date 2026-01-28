package main

import (
    "belajar-gin/routes" // Pastikan ini benar
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    routes.Routes(r) // Memanggil fungsi yang baru saja Anda buat
    r.Run()
}