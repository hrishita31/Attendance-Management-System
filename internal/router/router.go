package router

import (
	"attendance-system/internal/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.POST("/newUser", handler.NewUser())
	e.GET("/validateUser", handler.ValidateUser())
	e.GET("/checkUsername", handler.CheckUsername())
	//e.GET("/validateCreds", handler.ValidateCreds())

	e.POST("/enterAttendance", handler.EnterAttendance())
	e.GET("/getAttendanceByDate", handler.CheckAttendanceByDate())
	e.GET("/getTotalAttendance", handler.TotalAttendance())

	e.POST("/generateStudentID", handler.GenerateStudentID())
	return e

}
