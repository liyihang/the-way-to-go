package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

func (person Person1) getNameAndAge() (string, int) {
	return person.name, person.age
}

type Student1 struct {
	Person1
	speciality string
}

func (student Student1) getSpeciality() string {
	return student.speciality
}
func main() {
	student := new(Student1)
	student.name = "zhangsan"
	student.age = 18
	name, age := student.getNameAndAge()
	//speciality := student.getSpeciality()
	fmt.Print(name, age)
}
