package main
import(
	"net/http"
	"github.com/labstack/echo"
	"aooing.com-echo/crontrollers/user"
	"aooing.com-echo/crontrollers/file"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template)Render(w io.Writer, name string, data interface{}, c echo.Context)error{
	return t.templates.ExecuteTemplate(w, name, data)
}

func Routers(e *echo.Echo){
	e.Static("/media/*", "data")
	e.Static("/static/*", "static")
	e.GET("/", func(c echo.Context)error {
		return c.Render(http.StatusOK, "hello.tpl", "")
	})

	filegroup := e.Group("/file")
	file.RegisterFile(filegroup)

	api := e.Group("/api").Group("/v1")
	user.Regisger(api)
}
func PreCompileTpls(e *echo.Echo){
	t := &Template{
		templates: template.Must(template.ParseGlob("tpls/*.tpl")),
	}
	e.Renderer = t
}

func main(){
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	Routers(e)
	PreCompileTpls(e)
	e.Logger.Fatal(e.Start(":9091"))
}