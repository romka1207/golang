package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

type Note struct {
	Name    string
	Surname string
	Note    string
}

var sl []Note

func main() {
	e := echo.New()
	e.POST("/save_note", SaveNote)
	e.GET("/get_notes", GetNotes)
	log.Fatalln(e.Start("127.0.0.1:4000"))
}

func SaveNote(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}

	n := Note{}

	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}

	sl = append(sl, n)

	fmt.Println("Имя:", n.Name)
	fmt.Println("Фамилия", n.Surname)
	fmt.Println("Заметка", n.Note)
	return c.NoContent(200)
}

func GetNotes(c echo.Context) error {
	return c.JSON(http.StatusOK, &sl)
}
