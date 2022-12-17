package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// type httpHandler struct{}
type Note struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Note    string `json:"note,omitempty"`
}

func getHello(c echo.Context) error {
	params := c.QueryParams()["q"]
	if len(params) != 0 {
		fmt.Println("P", params[0])
	} else {
		fmt.Println("NO P")
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func postCreateNote(c echo.Context) error {
	defer mutex.Unlock()
	mutex.Lock()
	body := c.Request().Body
	decoder := json.NewDecoder(body)
	var note Note
	err := decoder.Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	notes = append(notes, note)
	log.Println("Name: " + note.Name + " Surname: " + note.Surname + " Text:" + note.Note)
	return c.String(http.StatusOK, "OK")

}

func getReadNote(c echo.Context) error {
	defer mutex.Unlock()
	mutex.Lock()
	params := c.QueryParams()["id"]
	txt := "err"
	if len(params) != 0 {
		txt = params[0]
	}
	NoteID, err := strconv.ParseInt(txt, 10, 64)
	if err == nil {
		NoteID -= 1
		if NoteID >= 0 && NoteID < int64(len(notes)) {
			note := notes[NoteID]
			jsonValue, _ := json.Marshal(note)
			return c.String(http.StatusOK, string(jsonValue))
		}

	}
	return c.String(http.StatusInternalServerError, "err")
}

func getNotesCount(c echo.Context) error {
	defer mutex.Unlock()
	mutex.Lock()
	l := strconv.Itoa(len(notes))
	return c.String(http.StatusOK, l)
}

var notes []Note
var mutex sync.Mutex

func main() {
	//http.HandleFunc("/createNote", getCreateNote)
	//http.HandleFunc("/getNoteCount", getNotesCount)
	//http.HandleFunc("/readNote", getReadNote)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", getHello)
	e.GET("/getNoteCount", getNotesCount)
	e.POST("/createNote", postCreateNote)
	e.GET("readNote", getReadNote)

	e.Logger.Fatal(e.Start(":1323"))
}
