package main

import (
	"fmt"
	"list/storages/list"
	"list/storages/model"
	"list/storages/slice"
	"log"
)

type People struct {
	ID   int
	Name string
}

func main() {
	l := list.NewList()
	_, err := model.Add(l, People{ID: 4, Name: "a"})
	if err != nil {
		log.Fatalln(err)
	}
	model.Add(l, People{ID: 3, Name: "c"})
	model.Add(l, People{ID: 2, Name: "d"})
	model.Add(l, People{ID: 1, Name: "b"})
	model.Print(l)
	model.Delete(l, 3)
	fmt.Println("Sort List")
	model.Sort(l, func(i, j any) bool {
		v1, ok := i.(People)
		if !ok {
			log.Fatalln("v1, ok := i.(People)")
		}
		v2, ok := j.(People)
		if !ok {
			log.Fatalln("v2, ok := j.(People)")
		}
		return v1.ID > v2.ID
	})
	model.Print(l)
	//fmt.Println(model.Get(l, 1))

	s := slice.Slice{}
	model.Add(&s, 15)
	model.Add(&s, 9)
	model.Add(&s, -2)
	model.Add(&s, 35)
	model.Add(&s, 7)
	model.Add(&s, 1)
	model.Print(&s)
	model.Delete(&s, 3)
	fmt.Println("Sort List")
	model.Sort(&s, func(i, j any) bool {
		v1, ok := i.(int)
		if !ok {
			log.Fatalln("v1, ok := i.(People)")
		}
		v2, ok := j.(int)
		if !ok {
			log.Fatalln("v2, ok := j.(People)")
		}
		return v1 > v2
	})
	s.Print()
}
