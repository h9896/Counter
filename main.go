package main

import (
	"log"
	"net/http"
	"os"

	con "github.com/h9896/Counter/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var limit = 61
var port = "80"

func main() {
	port = os.Getenv("PORT")
	e := EchoEngine(limit, nil)
	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

// EchoEngine - echo setting
func EchoEngine(interval int, ipList []string) *echo.Echo {
	// Echo instance
	e := echo.New()
	counter := &state{cal: con.NewCounter(interval), internalIP: ipList}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(counter.permission())

	// Route => handler
	e.GET("/", counter.getRequest)
	e.GET("/all", counter.getAllRequest)
	return e
}
func (s *state) permission() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, ip := range s.internalIP {
				if ip == c.RealIP() {
					return next(c)
				}
			}
			allow, err := s.cal.GetPermission(c.RealIP(), 61)
			if err != nil {
				log.Println(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Error, ", err)
			}
			if allow {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusTooManyRequests, "Error, request limit exceeded")
		}
	}
}

func (s *state) getRequest(c echo.Context) error {
	val := s.cal.GetNumber(c.RealIP())
	return c.JSON(http.StatusOK, echo.Map{
		"requestNum": val,
	})
}
func (s *state) getAllRequest(c echo.Context) error {
	return c.JSON(http.StatusOK, s.cal.GetAllNumber())
}

type (
	state struct {
		cal        *con.Counter
		internalIP []string
	}
)
