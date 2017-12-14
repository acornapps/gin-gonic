package main

import (

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	services := getAllServiceList()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": services}, "index.html")
}
