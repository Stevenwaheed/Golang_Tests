package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Student struct{
	Id int
	student_name string
}

/*
	This function responsible for displaying the Id and name of each student.
*/
func display(st []Student){
	fmt.Println("+--------+------------------+")
	fmt.Println("|   id   |      student     |")
	fmt.Println("+--------+------------------+")

	for _, student := range st{
		fmt.Println("   ", student.Id, "      ", student.student_name, "")
	}

	fmt.Println("+--------+------------------+")
}

/*
	This function responsible for swap the students.
	if the number of the students are odd then it leaves the last one without swaping.
*/
func swap(st []Student) []Student{
	if len(st) % 2 == 0{
		for student := 0; student <= len(st) - 2; student += 2{
			st[student].student_name, st[student + 1].student_name = st[student + 1].student_name, st[student].student_name
		}
	} else {
		for student := 0; student <= len(st) - 3; student += 2{
			st[student].student_name, st[student+1].student_name = st[student+1].student_name, st[student].student_name
		}
	}
	
	return st
}


func main(){
	
	fmt.Println("Enter the number of students: ")
	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')        // take the number of students from the user
	numberOfStudents, _ := strconv.Atoi(num[:len(num)-2])

	var student Student
	var students_list []Student
	var id = 0

	for i:=0; i<numberOfStudents; i++{
		fmt.Println("Enter the student name: ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')       // take each student name from the user
		id += 1

		student.Id = id
		student.student_name = name
		students_list = append(students_list, student)     // store each student in an array
	}

	new_students_list := swap(students_list)

	display(new_students_list)
}