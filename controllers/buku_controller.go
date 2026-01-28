package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Daftar Buku"})
}
