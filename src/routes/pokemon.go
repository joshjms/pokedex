package routes

import (
	"pokedex/src/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Init(db *[]models.Pokemon) *echo.Echo {
	e := echo.New()

	g := e.Group("/api/v1")

	g.GET("/pokemon", func(c echo.Context) error {
		return c.JSON(200, db)
	})

	// g.GET("/pokemon/image/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	idInt, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		return c.JSON(400, map[string]string{
	// 			"message": "Invalid ID",
	// 		})
	// 	}

	// 	for _, pokemon := range *db {
	// 		if pokemon.ID == idInt {
	// 			imageURL := pokemon.Image.Hires

	// 			response, err := http.Get(imageURL)
	// 			if err != nil {
	// 				return c.String(http.StatusInternalServerError, "Failed to fetch image")
	// 			}
	// 			defer response.Body.Close()

	// 			contentType := response.Header.Get("Content-Type")
	// 			c.Response().Header().Set("Content-Type", contentType)

	// 			_, err = io.Copy(c.Response().Writer, response.Body)
	// 			if err != nil {
	// 				return c.String(http.StatusInternalServerError, "Failed to copy image data")
	// 			}

	// 			return nil
	// 		}
	// 	}
	// 	return c.JSON(404, map[string]string{
	// 		"message": "Pokemon not found",
	// 	})
	// })

	g.GET("/pokemon/:id", func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(400, map[string]string{
				"message": "Invalid ID",
			})
		}

		for _, pokemon := range *db {
			if pokemon.ID == idInt {
				return c.JSON(200, pokemon)
			}
		}
		return c.JSON(404, map[string]string{
			"message": "Pokemon not found",
		})
	})

	return e
}
