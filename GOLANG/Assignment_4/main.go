package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Let's create our Car struct. This will contain information about a car

// Car contains information about a single car
type Car struct {
	ID   int    `json:"CarID" binding:"required"`
	Rate int    `json:"Car Rating"`
	Car  string `json:"Car Name" binding:"required"`
}

// We'll create a list of Vehicles
var cars = []Car{
	Car{1, 0, "Nissan Magnite"},
	Car{2, 0, "Kia Sonet"},
	Car{3, 0, "Maruti Swift"},
	Car{4, 0, "Hyundai Creta"},
	Car{5, 0, "Tata Altroz"},
	Car{6, 0, "Renault KWID"},
	Car{7, 0, "Mahindra XUV300"},
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	// Our API will consit of just two routes
	// /cars - which will retrieve a list of cars a user can see
	// /cars/rate/:carID - which will capture rating sent to a particular car
	api.GET("/cars", CarHandler)
	api.POST("/cars/rate/:carID", CarRating)

	// Start and run the server
	router.Run(":8000")
}

// CarHandler retrieves a list of available cars
func CarHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, cars)
}

// carModel gives the year of particular car
func CarRating(c *gin.Context) {
	// confirm Joke ID sent is valid
	if carid, err := strconv.Atoi(c.Param("carID")); err == nil {
		// find car, and increment rating
		for i := 0; i < len(cars); i++ {
			if cars[i].ID == carid {
				cars[i].Rate += 1
			}
		}

		// return a pointer to the updated cars list
		c.JSON(http.StatusOK, &cars)
	} else {
		// Car ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}
