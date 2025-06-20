package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/go-singleflight/utils"
)

func fetchUser(id string) (map[string]interface{}, error) {
	fmt.Println("👤 Fetching User...", id)
	time.Sleep(500 * time.Millisecond)
	return map[string]interface{}{"id": id, "name": "Alice"}, nil
}

func fetchProfile(id string) (map[string]interface{}, error) {
	fmt.Println("📄 Fetching Profile...", id)
	time.Sleep(500 * time.Millisecond)
	return map[string]interface{}{"age": 30, "location": "Jakarta"}, nil
}

func fetchSettings(id string) (map[string]interface{}, error) {
	fmt.Println("⚙️ Fetching Settings...", id)
	time.Sleep(500 * time.Millisecond)
	return map[string]interface{}{"theme": "dark", "lang": "en"}, nil
}

// Result struct using TripleResult
type Result struct {
	User     map[string]interface{} `json:"user"`
	Profile  map[string]interface{} `json:"profile"`
	Settings map[string]interface{} `json:"settings"`
}

func main() {
	e := echo.New()

	e.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")

		key := utils.GenKey(id, "lalla")

		sf := &utils.Singleflight[Result]{Group: &utils.Group, Key: key}

		result, err := sf.ProcessWrapper(func() (Result, error) {
			user, err := fetchUser(id)
			if err != nil {
				return Result{}, err
			}
			profile, err := fetchProfile(id)
			if err != nil {
				return Result{}, err
			}
			settings, err := fetchSettings(id)
			if err != nil {
				return Result{}, err
			}
			return Result{
				User:     user,
				Profile:  profile,
				Settings: settings,
			}, nil
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, result)
	})

	fmt.Println("🚀 Server running at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
