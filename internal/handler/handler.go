package handler

import (
	"attendance-system/internal/infra"
	"attendance-system/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewUser() echo.HandlerFunc {
	client := infra.GetClient()

	return func(c echo.Context) error {
		name := c.QueryParam("Name")
		username := c.QueryParam("Username")
		password := c.QueryParam("Password")
		data, err := service.NewUser(name, username, password, client)

		if err != nil {
			return c.String(http.StatusInternalServerError, "incorrect")
		}

		return c.String(200, data)
	}
}

func ValidateUser() echo.HandlerFunc {
	client := infra.GetClient()

	return func(c echo.Context) error {
		name := c.QueryParam("Name")
		username := c.QueryParam("Username")
		password := c.QueryParam("Password")
		err := service.ValidateUser(name, username, password, client)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusUnauthorized, "user credentials invalid")
		}
		return c.String(http.StatusAccepted, "user credentials valid")

	}
}

func CheckUsername() echo.HandlerFunc {
	client := infra.GetClient()

	return func(c echo.Context) error {
		name := c.QueryParam("Name")
		username := c.QueryParam("Username")
		err := service.CheckUsername(name, username, client)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusUnauthorized, "user account created")
		}
		return c.String(http.StatusAccepted, "username already exists")
	}

}

func EnterAttendance() echo.HandlerFunc {
	client := infra.GetClient()

	return func(c echo.Context) error {
		date := c.QueryParam("Date")
		username := c.QueryParam("Username")
		cc := c.QueryParam("CC")
		gt := c.QueryParam("GT")
		dm := c.QueryParam("DM")
		cc1, err := strconv.ParseBool(cc)
		if err != nil {
			fmt.Println(err)
		}
		gt1, err := strconv.ParseBool(gt)
		if err != nil {
			fmt.Println(err)
		}
		dm1, err := strconv.ParseBool(dm)
		if err != nil {
			fmt.Println(err)
		}
		att, err := service.EnterAttendance(date, username, cc1, gt1, dm1, client)
		if err != nil {
			return c.String(http.StatusInternalServerError, "incorrect")
		}
		return c.String(http.StatusAccepted, att)
	}

}

func CheckAttendanceByDate() echo.HandlerFunc {
	client := infra.GetClient()

	return func(c echo.Context) error {
		date := c.QueryParam("Date")
		username := c.QueryParam("Username")
		attendance, err := service.CheckAttendanceByDate(date, username, client)
		if err != nil {
			return c.String(http.StatusUnauthorized, "no record exists for the given date and username")
		}
		return c.JSON(http.StatusAccepted, attendance)
	}
}

func TotalAttendance() echo.HandlerFunc {
	client := infra.GetClient()

	return func(c echo.Context) error {
		username := c.QueryParam("Username")
		subject := c.QueryParam("Subject")

		totalAtt, err := service.TotalAttendance(username, subject, client)
		if err != nil {
			return c.String(http.StatusUnauthorized, "attendance not found")
		}
		return c.JSON(http.StatusAccepted, totalAtt)
	}
}

func GenerateStudentID() echo.HandlerFunc {
	client := infra.GetClient()
	return func(c echo.Context) error {
		year := c.QueryParam("Year")
		year1, err := strconv.Atoi(year)
		if err != nil {
			fmt.Println(err)
		}
		dept := c.QueryParam("Department")
		newId, err := service.GenerateStudentID(year1, dept, client)
		if err != nil {
			return c.String(http.StatusUnauthorized, "id not created")
		}
		return c.JSON(http.StatusAccepted, newId)
	}
}

// func ValidateCreds() echo.HandlerFunc {
// 	client := infra.GetClient()

// 	return func(c echo.Context) error {
// 		username := c.QueryParam("Username")
// 		password := c.QueryParam("Password")
// 		err := service.ValidateCreds(username, password, client)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		return c.String(http.StatusAccepted, "username and passwords match")
// 	}
// }
