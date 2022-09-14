package controllers

import (
	"blog/connection"
	"blog/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateEducation(c echo.Context) error {
	education := new(models.Education)
	if err := c.Bind(education); err != nil {
		return err
	}
	education.Created_At = time.Now()
	education.Updated_At = time.Now()
	sqlStatement := "INSERT INTO education(degree, institute, subject, location, website_link, joining_date, leaving_date, description, created_at, updated_at, reg_no)VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	_, err := connection.DB.Query(sqlStatement, education.Degree, education.Institute, education.Subject, education.Location, education.Website_Link, education.Joining_Date, education.Leaving_Date, education.Description, education.Created_At, education.Updated_At, education.Reg_No)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, education)
}

func GetEducationInfo(c echo.Context) error {
	var educations []models.Education
	sqlStatement := "SELECT id, degree, institute, subject, location, website_link, joining_date, leaving_date, description, reg_no FROM education"
	rows, err := connection.DB.Query(sqlStatement)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var education models.Education
		err := rows.Scan(&education.Id, &education.Degree, &education.Institute, &education.Subject, &education.Location, &education.Website_Link,
			&education.Joining_Date, &education.Leaving_Date, &education.Description, &education.Reg_No)
		if err != nil {
			return err
		}
		educations = append(educations, education)
	}
	return c.JSON(http.StatusOK, educations)
}
