package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type QuizDataStruct struct {
	RoomID      string
	QuizID      string
	QuizText    string
	QuizAnswer  string
	QuizTimeout string
	Choice1     string
	Choice2     string
	Choice3     string
	Choice4     string
	NextQuiz    string
}

type Room struct {
	Id       int    `json:"id" param:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Author   string `json:author`
	Date     string `json:date`
	Comment  string `json:comment`
}

func StartQuiz(c echo.Context) error {
	// TODO: ここらへんDBと連携する
	QuizData := QuizDataStruct{
		RoomID:      c.Param("RoomID"),
		QuizID:      c.Param("QuizID"),
		QuizText:    "「家康の康ですよね？」は高専何年生のとき？",
		QuizAnswer:  "高専2年",
		QuizTimeout: "6000",
		Choice1:     "高専1年",
		Choice2:     "高専2年",
		Choice3:     "高専3年",
		Choice4:     "高専4年",
		NextQuiz:    "finish",
	}

	return c.Render(http.StatusOK, "quiz", QuizData)
}

func GetQuiz(c echo.Context) error {
	db := sqlConnect()
	room := Room{}
	if err := c.Bind(&room); err != nil {
		return err
	}
	id := c.Param("RoomID")
	db.Raw("select * from rooms where id=" + id).Scan(&room)
	defer db.Close()
	fmt.Println(room)
	return c.Render(http.StatusOK, "room", room)
}

func GetRoom(c echo.Context) error {
	db := sqlConnect()
	rooms := []Room{}
	db.Find(&rooms)
	defer db.Close()
	fmt.Println(rooms)
	return c.Render(http.StatusOK, "lobby", rooms)
}

func main() {
	db := sqlConnect()
	db.AutoMigrate(&User{})
	defer db.Close()

	e := echo.New()
	tl, err := template.New("t").ParseGlob("views/*.html")
	tl.ParseGlob("views/common/*.html")
	tl.ParseGlob("views/quiz/*.html")
	t := &Template{
		templates: template.Must(tl, err),
	}
	e.Renderer = t

	e.GET("/login", func(c echo.Context) error {
		data := struct{ info string }{"login"}
		return c.Render(http.StatusOK, "login", data)
	})
	e.GET("/home", func(c echo.Context) error {
		data := struct{ info string }{"home"}
		return c.Render(http.StatusOK, "home", data)
	})

	e.GET("/lobby/:QuizID", GetQuiz)
	e.GET("/lobby", GetRoom)
	e.GET("/room/:RoomID", GetQuiz)
	e.GET("/quiz/:RoomID/:QuizID", StartQuiz)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Here is root :)")
	})

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret1" {
			return true, nil
		}
		return false, nil
	}))

	e.Logger.Fatal(e.Start(":8080"))

}

func sqlConnect() (database *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	DBMS := "mysql"
	PROTOCOL := "tcp(db:3306)"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	fmt.Println("user : " + USER)
	fmt.Println("password : " + PASS)
	fmt.Println("DB name : " + DBNAME)
	fmt.Println("DBMS : " + DBMS)
	fmt.Println("protocol : " + PROTOCOL)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("DB接続失敗🤪🤪🤪")
	} else {
		fmt.Println("やったね😊")
	}
	//db.LogMode(true)
	return db
}
