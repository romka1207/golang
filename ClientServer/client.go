package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Note struct {
	Name    string
	Surname string
	Note    string
}

var httpClient = &http.Client{}

func main() {
	v := Input()
	v.PostNote()

	execute := true
	for execute == true {
		execute = Choose()
	}
}

func Choose() bool {
	fmt.Printf("\nЧто делаем дальше?\n")
	fmt.Printf("\ny - продолжить, n - завершить, p - показать всё\n")

	var choice string
	_, err := fmt.Scanf("%s\n", &choice)
	if err != nil {
		return false
	}

	if choice == "n" {
		return false
	}

	if choice == "y" {
		v := Input()
		v.PostNote()
		return true
	}

	if choice == "p" {
		PrintNotes()
		return true
	}

	CallClear()
	fmt.Printf("Пожалуйста, выберите один из предложенных вариантов\n")
	return true
}

func Input() Note {
	CallClear()

	myScan := bufio.NewScanner(os.Stdin)

	fmt.Println("Имя")
	myScan.Scan()
	name := myScan.Text()

	fmt.Println("Фамилия")
	myScan.Scan()
	surname := myScan.Text()

	fmt.Println("Заметка")
	myScan.Scan()
	note := myScan.Text()

	v := Note{name, surname, note}

	return v
}

func (v *Note) PostNote() {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	bb := bytes.Buffer{}
	bb.Write(marshal)

	req, err := http.NewRequest("POST", "http://127.0.0.1:4000/save_note", &bb)
	if err != nil {
		log.Fatal(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 200 {
		fmt.Printf("\nВведённые вами данные:")
		fmt.Printf("\nИмя - %s, Фамилия - %s, Заметка - %s\n", v.Name, v.Surname, v.Note)
	} else {
		log.Fatal(err)
	}
}

func PrintNotes() {
	var notes []Note

	res, err := http.Get("http://127.0.0.1:4000/get_notes")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if json.Unmarshal(body, &notes) != nil {
		log.Fatal(err)
		return
	}

	CallClear()
	for index, note := range notes {
		fmt.Printf("Заметка № %d\n", index+1)
		fmt.Printf("Имя - %s\nФамилия - %s\nЗаметка - %s\n\n", note.Name, note.Surname, note.Note)
	}
}

func CallClear() {
	fmt.Print("\033[H\033[2J")
}
