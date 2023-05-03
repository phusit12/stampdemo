package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/phusit12/stampdemo.git/models"
)

func NewUserRoute(uGroup *echo.Group) {

	uGroup.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		u := models.GetUser(id)
		if u == nil {
			return c.JSON(http.StatusBadRequest, models.MessageUserRes{
				Message: "User not found",
			})
		}
		return c.JSON(http.StatusOK, models.GetUserRes{
			Id:   u.GetID(),
			Name: u.GetName(),
			Date: time.Now().UTC(),
		})
	})
	uGroup.POST("/register", func(c echo.Context) error {
		req := new(models.GetUserRes)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, models.MessageUserRes{
				Message: err.Error(),
			})
		}
		id := req.Id
		name := req.Name
		if id == "" || name == "" {
			return c.JSON(http.StatusBadRequest, models.MessageUserRes{
				Message: "Invalid id or name",
			})
		}

		if models.GetUser(id) != nil {
			return c.JSON(http.StatusBadRequest, models.MessageUserRes{
				Message: "This id registerd",
			})
		}
		models.NewUser(id, name)
		return c.JSON(http.StatusOK, models.MessageUserRes{
			Message: "Register success",
		})
	})

	uGroup.GET("/:id/name", func(c echo.Context) error {
		id := c.Param("id")
		u := models.GetUser(id)
		if u == nil {
			return c.JSON(http.StatusBadRequest, models.MessageUserRes{
				Message: "User not found",
			})
		}
		return c.JSON(http.StatusOK, u.GetName())
	})
}

//uGroup.GET("/id", func(c echo.Context) error {
//return c.JSON(http.StatusOK, u.GetID())
//})

//e.logger erorrจาก echo e.start path เครื่องเราและเอาไว้ดูว่ามีปัญหาการรันไม่ได้
