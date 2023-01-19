package routing

import (
	"fmt"

	"github.com/ksaepudin/movie/src/routing/group"
	"github.com/labstack/echo/v4"
)

// this layer for routing
// This func is only for group routing handles that are listed in the Route func block
func Route() {
	// add func echo new
	e := echo.New()

	// Registerd for service routing group
	group.MovieRoute(e)

	// this set port for run apps
	port := fmt.Sprintf(":%s", "9999")

	// this for run apps
	e.Logger.Fatal(e.Start(port))
}
