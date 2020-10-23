package main

import "fmt"

var allStudents = []student{}

type student struct {
	name string
	age  int
	id   int
}

var name string
var age int

func main() {

	var value int
	id := 0

menu:
	fmt.Println("\n WELCOME STUDENTS")
	fmt.Println("Please select a suitable option from below:-")
	fmt.Println("1--->Add a new student")
	fmt.Println("2--->Edit details of any student")
	fmt.Println("3--->Display details of any student")
	fmt.Println("4--->Delete any student")
	fmt.Println("5--->Exit")
	fmt.Println("(Default:-)List all")
	fmt.Scanln(&value)

	switch {
	case value == 1:
		{
			fmt.Println("Enter the Name of student")
			fmt.Scanln(&name)
			fmt.Println("Enter the Age of student")
			fmt.Scanln(&age)

			id = id + 1
			newStudent := student{name: name, age: age, id: id}
			allStudents = append(allStudents, newStudent)

			goto menu
		}
	case value == 2:
		{
			fmt.Println("Please, enter the id of the student")
			fmt.Scanln(&id)
			var name2 string
			var age2 int
			for i, v := range allStudents {
				if id == v.id {
					fmt.Println("Enter the Name of student")
					fmt.Scanln(&name2)
					fmt.Println("Enter the Age of student")
					fmt.Scanln(&age2)
					v.name = name2
					v.age = age2
					allStudents[i] = v
				}
			}
			goto menu
		}
	case value == 3:
		{
			{
				fmt.Println("Please, Enter the id ")
				fmt.Scanln(&id)
				for _, v := range allStudents {
					if id == v.id {
						fmt.Println("\n OUR SEARCH RESULT IS:- \n", v)
					}
				}
			}
			goto menu
		}
	case value == 4:
		{
			fmt.Println("Enter the id of the record to be deleted")
			fmt.Scanln(&id)
			for i, v := range allStudents {
				if id == v.id {
					fmt.Println("Student with these details has been deleted ---> ", v)
					allStudents = append(allStudents[:i], allStudents[i+1:]...)
				}
			}
			goto menu
		}
	case value == 5:
		break
	default:
		{
			fmt.Println("\n Listing all")
			fmt.Println(allStudents)
			goto menu
		}
	}
}
