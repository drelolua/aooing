package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"aooing.com-echo/models"
	"golang.org/x/crypto/bcrypt"
)

// User 用户
type User struct {
}

// Regisger 注册
func Regisger(g *echo.Group) {
	u := new(User)
	group := g.Group("/user")
	{
		group.GET("/show", u.Show)
		group.GET("/:name", u.GetUser)
		group.GET("/insert", u.Insert)
		group.GET("/all", u.All)
		group.GET("/reg", u.Register)
		group.GET("/login", u.Login)
		group.GET("/list", u.Userp)
	}
}

// Insert 插入
func (u *User) Insert(c echo.Context) error {
	um := new(models.User)
	um.Name = "我擦"
	um.Age = 23
	if err := um.Insert(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, um)
}

// All 获取全部用户
func (u *User) All(c echo.Context) error {
	um := new(models.User)
	uml := new([]models.User)
	_ = um.GetData(uml)

	return c.JSON(http.StatusOK, uml)
}

// GetUser 获取用户
//e.GET("/user/:id", user.GetUser)
func (u *User) GetUser(c echo.Context) error {
	name := c.Param("name")
	um := models.User{Name: name}
	type Au struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Message string `json:"message"`
		Phone   string `json:"phone"`
	}
	au := new(Au)
	um.Find(um.CName(), um, nil).One(&au)

	return c.JSON(http.StatusOK, au)
}

// Show 展示
//e.GET("/user/show", Show)
/*
	根据传入的姓名和年龄，返回json格式的用户信息。
*/
func (u *User) Show(c echo.Context) error {
	name := c.QueryParam("name")
	ages := c.QueryParam("age")
	c.Set("content-type", "application/json")
	type JUser struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	age, _ := strconv.Atoi(ages)
	var ju = JUser{Name: name, Age: age}
	s := c.JSON(http.StatusOK, ju)
	return s
}

// Register 注册
// 生成密码 GET /reg
/*
	使用bcrypt计算密码，每次生成不同的密码，但是数据库只需保存一份
*/
func (u *User) Register(c echo.Context) error {
	userName := c.QueryParam("name")
	passwd := c.QueryParam("passwd")
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		//
	}
	type up struct {
		UserName  string `json:"user_name"`
		PassWord  string `json:"pass_word"`
		EncPasswd string `json:"enc_passwd"`
	}
	encodePwd := string(hash)

	rup := up{UserName: userName, PassWord: passwd, EncPasswd: encodePwd}

	return c.JSON(http.StatusOK, rup)

}

// Userp 分页
func (u *User) Userp(c echo.Context) error {
	pn, _ := strconv.Atoi(c.QueryParam("pn")) // 第几页
	n, _ := strconv.Atoi(c.QueryParam("n"))   // 每页多少个

	pu := models.User{}
	ul := new([]models.User)
	pu.GetPageData(((pn - 1) * n), n, ul)
	return c.JSON(http.StatusOK, ul)
}

// Login 登录
func (u *User) Login(c echo.Context) error {
	ckPasswd := c.QueryParam("ck")
	encodePwd := c.QueryParam("en")
	type Rp struct {
		Ok bool `json:"ok"`
	}
	res := new(Rp)
	err := bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(ckPasswd))
	if err != nil {
		res.Ok = false
	} else {
		res.Ok = true
	}
	return c.JSON(http.StatusOK, res)

}
