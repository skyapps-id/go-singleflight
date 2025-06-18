package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

// Dummy DB fetch function
func fetchUserFromDB(id string) (map[string]interface{}, error) {
	// Simulate database delay
	fmt.Printf("⏳ [DB] Fetching user ID: %s\n", id)
	time.Sleep(2 * time.Second) // simulate slow DB
	// Simulate a user record
	user := map[string]interface{}{
		"id":   id,
		"name": "John Doe",
	}
	fmt.Printf("✅ [DB] Finished fetching user ID: %s\n", id)
	return user, nil
}

func main() {
	e := echo.New()

	e.GET("/singleflight/user/:id", func(c echo.Context) error {
		id := c.Param("id")

		// Use singleflight to deduplicate calls for the same ID
		v, err, shared := group.Do(id, func() (interface{}, error) {
			return fetchUserFromDB(id)
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		if shared {
			fmt.Println("📦 [Shared] Returned cached result for ID:", id)
		}

		return c.JSON(http.StatusOK, v)
	})

	e.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")

		resp, _ := fetchUserFromDB(id)

		return c.JSON(http.StatusOK, resp)
	})

	fmt.Println("🚀 Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
