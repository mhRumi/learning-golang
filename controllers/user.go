package controllers

import (
	"blog/connection"
	"blog/models"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sethvargo/go-password/password"

	"github.com/labstack/echo/v4"
)

func Registration(c echo.Context) error {
	u := new(models.User)
	email := c.FormValue("email")
	if err := c.Bind(u); err != nil {
		return err
	}
	sqlStatement := "INSERT INTO users (name, reg_no, batch, createdat, updatedat) VALUES ($1, $2, $3, $4, $5)"
	_, err := connection.DB.Query(sqlStatement, u.Name, u.Reg_No, u.Batch, time.Now(), time.Now())
	if err != nil {
		fmt.Println(err)
		return err
	}
	password, err := password.Generate(8, 4, 3, false, false)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Password: " + password)
	hashedpassword, _ := HashPassword(password)
	sqlStatement = "INSERT INTO credentials (reg_no, email, password, createdat, updatedat) VALUES($1, $2, $3, $4, $5)"
	_, err = connection.DB.Query(sqlStatement, u.Reg_No, email, hashedpassword, time.Now(), time.Now())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func GetUserRegNo(c echo.Context) float64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	reg_no := claims["reg_no"].(float64)
	return reg_no
}
