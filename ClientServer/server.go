package main

import (
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/save_note", SaveNote)
	http.HandleFunc("/get_notes", GetNotes)
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", nil))
}

func SaveNote(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	n := Note{}

	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Println(err)
	}

	sl = append(sl, n)

	fmt.Println("Имя:", n.Name)
	fmt.Println("Фамилия", n.Surname)
	fmt.Println("Заметка", n.Note)
}

func GetNotes(rw http.ResponseWriter, req *http.Request) {
	resp, err := json.Marshal(sl)
	if err != nil {
		log.Fatal(err)
	}

	_, err = rw.Write(resp)
	if err != nil {
		return
	}
}
