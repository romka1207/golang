package main

import (
	"bufio"
	"fmt"
	"os"T
)

type Note struct {
	name    string
	surname string
	note    string
}

func main() {
	var sl []Note

	v := input()
	sl = append(sl, v)

	fl := true
	for fl == true {
		fl = choose(&sl)
	}
}

func input() Note {
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

	fmt.Printf("\nВведённые вами данные:")
	fmt.Printf("\nИмя - %s, Фамилия - %s, Заметка - %s\n", v.name, v.surname, v.note)

	return v
}

func choose(sl *[]Note) bool {

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
		v := input()
		*sl = append(*sl, v)
		return true
	}

	if choice == "p" {
		CallClear()
		for i := 0; i < len(*sl); i++ {
			fmt.Printf("Заметка № %d\n", i+1)
			fmt.Printf("Имя - %s\nФамилия - %s\nЗаметка - %s\n\n", (*sl)[i].name, (*sl)[i].surname, (*sl)[i].note)
		}
		return true
	}

	CallClear()
	fmt.Printf("Пожалуйста, выберите один из предложенных вариантов\n")
	return true
}

func CallClear() {
	fmt.Print("\033[H\033[2J")
}
