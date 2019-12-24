package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	studentsAges := [2]int{29, 30}
	studentsStruct := []Student{
		Student{
			name: "hilder",
			age:  29,
		},
		Student{
			name: "hilder",
			age:  29,
		},
	}
	fmt.Println("Inicializando studiante", studentsAges)
	fmt.Println("Inicializando studiante Struct", studentsStruct)
}
func calculateAges(student1 Student, student2 Student) int {
	return student1.age + student2.age
}
func avgAges(student1 Student, student2 Student) int {
	return (student1.age + student2.age) / 2
}
