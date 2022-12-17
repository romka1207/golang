package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Note    string `json:"note"`
}

func getCreateNote(w http.ResponseWriter, r *http.Request) {
	defer mutex.Unlock()
	mutex.Lock()
	decoder := json.NewDecoder(r.Body)
	var note Note
	err := decoder.Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	notes = append(notes, note)
	log.Println("Name: " + note.Name + " Surname: " + note.Surname + " Text:" + note.Note)
	fmt.Fprintf(w, "Name: "+note.Name+" Surname: "+note.Surname+" Text:"+note.Note)

}

/*
func getHelloPage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, "+name)
}
*/

func getReadNote(w http.ResponseWriter, r *http.Request) {
	defer mutex.Unlock()
	mutex.Lock()
	txt := r.URL.Query().Get("id")
	NoteID, err := strconv.ParseInt(txt, 10, 64)
	if err == nil {
		NoteID -= 1
		if NoteID >= 0 && NoteID < int64(len(notes)) {
			note := notes[NoteID]
			jsonValue, _ := json.Marshal(note)
			fmt.Fprintf(w, string(jsonValue))
		}

	}
	fmt.Fprintf(w, "err")
}

func getNotesCount(w http.ResponseWriter, r *http.Request) {
	defer mutex.Unlock()
	mutex.Lock()
	l := strconv.Itoa(len(notes))
	fmt.Fprintf(w, l)
}

var notes []Note
var mutex sync.Mutex

func main() {
	//http.HandleFunc("/hello", getHelloPage)
	http.HandleFunc("/createNote", getCreateNote)
	http.HandleFunc("/getNoteCount", getNotesCount)
	http.HandleFunc("/readNote", getReadNote)
	err := http.ListenAndServe(":4862", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
