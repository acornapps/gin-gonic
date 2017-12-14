package main

import
(
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
)


func initializeRoutes() {

	router.GET("/", showIndexPage)

	api1 := router.Group("/api")
	{

		api1.GET("/svc/:id", func(c *gin.Context) {

			id , err := strconv.Atoi(c.Param("id"))
			if err != nil {
				fmt.Printf("Error : %v ",err)
			}

			svc , err := getServiceByID(id)

			if err != nil {
				fmt.Printf("Error : %v ",err)
			}
			c.JSON(200, gin.H{
				"ID" : svc.ID,
				"Title": svc.Title,
				"Content": svc.Content,
			})
		})

	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})


}