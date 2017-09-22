package file

import (
	"github.com/labstack/echo"
	"os"
	"io"
	"net/http"
	"crypto/md5"
	"fmt"
)

const(
	DATA_PATH string = "data/"
	MEDIA_PREFIX = "media/"
)

func RegisterFile(g *echo.Group){
	g.GET("/upload", UploadGet)
	g.POST("/upload", UploadPost)
}

//e.GET("/file/upload", file.UploadGet)
func UploadGet(c echo.Context)error{
	return c.Render(http.StatusOK, "upload.tpl", "")
}

//e.POST("/file/upload", file.UploadPost)
func UploadPost(c echo.Context)error{
	host := c.Request().Host
	name := c.FormValue("name")
	email := c.FormValue("email")

	form, err := c.MultipartForm()
	if err != nil{
		return err
	}
	postFiles := form.File["files"]

	type img struct {
		Url string `json:"url"`
		Name string `json:"name"`
		Email string `json:"email"`
	}
	var resp []img

	for _, file := range postFiles{
		src, err := file.Open()
		if err != nil{
			return err
		}
		defer src.Close()

		md5_filename := fmt.Sprintf("%x",md5.Sum([]byte(file.Filename)))
		dst, err := os.Create(DATA_PATH+ md5_filename)
		if err != nil{
			return err
		}
		defer dst.Close()

		// Copy
		if _, err := io.Copy(dst, src);err != nil {
			return err
		}
		resp = append(resp, img{Url:"http://"+host+"/" +MEDIA_PREFIX+md5_filename, Email:email, Name:name})
	}
	return c.JSON(http.StatusOK, resp)
}