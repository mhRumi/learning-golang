package controllers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

func UploadImage(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	var fileNames []string
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(file.Size)
		if file.Size > 2000 {
			return errors.New("image size is huge, try to compress")
		}
		defer src.Close()
		filename := random.String(30) + file.Filename[strings.LastIndex(file.Filename, "."):]
		dst, err := os.Create("./assets/media/" + filename)
		fileNames = append(fileNames, filename)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}
	return c.JSON(http.StatusCreated, fileNames)
}
