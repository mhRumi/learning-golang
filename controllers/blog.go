package controllers

import (
	"blog/connection"
	"blog/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func PostBlog(c echo.Context) error {
	blog := new(models.Blog)
	reg_no := GetUserRegNo(c)
	if err := c.Bind(blog); err != nil {

		fmt.Println(err)
		return err
	}
	sqlStatement := "INSERT INTO blogs(title, content, image, createdAt, updatedAt, reg_no) VALUES($1, $2, $3, $4, $5, $6)"
	rows, err := connection.DB.Query(sqlStatement, blog.Title, blog.Content, pq.Array(blog.Image), time.Now(), time.Now(), reg_no)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rows)
}

func GetApprovedBlogs(c echo.Context) error {
	var blogs []models.Blog
	sqlStatement := "SELECT id, title, content, image, createdat, updatedat, reg_no FROM blogs where isapproved=TRUE"
	rows, err := connection.DB.Query(sqlStatement)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, pq.Array(&blog.Image), &blog.CreatedAt, &blog.UpdatedAt, &blog.Reg_No)
		if err != nil {
			return err
		}
		blogs = append(blogs, blog)
	}
	return c.JSON(http.StatusOK, blogs)
}

func GetPendingBlogs(c echo.Context) error {
	var blogs []models.Blog
	sqlStatement := "SELECT id, title, content, image, createdat, updatedat, reg_no FROM blogs where isapproved=FALSE"
	rows, err := connection.DB.Query(sqlStatement)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, pq.Array(&blog.Image), &blog.CreatedAt, &blog.UpdatedAt, &blog.Reg_No)
		if err != nil {
			return err
		}
		blogs = append(blogs, blog)
	}
	return c.JSON(http.StatusOK, blogs)
}

func UpdateBlog(c echo.Context) error {
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStatement := "UPDATE blogs SET title=$1, content=$2, image=$3, updatedat=$4, reg_no=$5 WHERE id=$6"
	_, err := connection.DB.Query(sqlStatement, blog.Title, blog.Content, pq.Array(blog.Image), time.Now(), blog.Reg_No, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, blog)
}
