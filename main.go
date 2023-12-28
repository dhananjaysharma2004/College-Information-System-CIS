package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type studentDetails struct {
	Name   string
	Course string
	Branch string
	Age    int
	CGPA   float32
}

type teacherDetails struct {
	Name     string
	Subject  string
	position string
	HomeCity string
}

type staffDetails struct {
	Name      string
	MaxSalary int
	HomeCity  string
}

type universityData struct {
	Students []studentDetails
	Teachers []teacherDetails
	Staff    []staffDetails
}

var data universityData
var mutex sync.Mutex

func loadData() error {
	file, err := ioutil.ReadFile("university.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func saveData() error {
	mutex.Lock()
	defer mutex.Unlock()
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("university.json", file, 0644)
}

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("WELCOME LADIES AND GENTLEMEN!!!")
	loadData()

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1 - Add Student")
		fmt.Println("2 - Add Teacher")
		fmt.Println("3 - Add Staff")
		fmt.Println("4 - View Details")
		fmt.Println("5 - Search Details")
		fmt.Println("6 - Exit")
		var choice int
		fmt.Scan(&choice)
		readLine()

		switch choice {
		case 1:
			var stu studentDetails
			fmt.Println("Enter NAME of the student:")
			stu.Name = readLine()
			fmt.Println("Enter COURSE of the student:")
			fmt.Scan(&stu.Course)
			fmt.Println("Enter BRANCH of the student:")
			fmt.Scan(&stu.Branch)
			fmt.Println("Enter AGE of the student:")
			fmt.Scan(&stu.Age)
			fmt.Println("Enter CGPA of the student:")
			fmt.Scan(&stu.CGPA)
			mutex.Lock()
			data.Students = append(data.Students, stu)
			mutex.Unlock()
			fmt.Println("________________________")
			fmt.Println(" THANK YOU!! Student details added.")
			fmt.Println("________________________")
			saveData()

		case 2:
			var tea teacherDetails
			fmt.Println("Enter NAME of the teacher:")
			tea.Name = readLine()
			fmt.Println("Enter SUBJECT of the teacher:")
			tea.Subject = readLine()
			fmt.Println("Enter POSITION of the teacher:")
			tea.position = readLine()
			fmt.Println("Enter HOMECITY of the teacher:")
			fmt.Scan(&tea.HomeCity)
			mutex.Lock()
			data.Teachers = append(data.Teachers, tea)
			mutex.Unlock()
			fmt.Println("________________________")
			fmt.Println(" THANK YOU!! Teacher details added.")
			fmt.Println("________________________")
			saveData()

		case 3:
			var sta staffDetails
			fmt.Println("Enter NAME of the staff:")
			sta.Name = readLine()
			fmt.Println("Enter MAXSALARY of the staff:")
			fmt.Scan(&sta.MaxSalary)
			fmt.Println("Enter HOMECITY of the staff:")
			fmt.Scan(&sta.HomeCity)
			mutex.Lock()
			data.Staff = append(data.Staff, sta)
			mutex.Unlock()
			fmt.Println("________________________")
			fmt.Println(" THANK YOU!! Staff details added.")
			fmt.Println("________________________")
			saveData()

		case 4:
			fmt.Println("Viewing University Details:")
			fmt.Println("Students:")
			for _, s := range data.Students {
				fmt.Printf("Name: %s\nCourse: %s\nBranch: %s\nAge: %d\nCGPA: %f\n", s.Name, s.Course, s.Branch, s.Age, s.CGPA)
			}
			fmt.Println("Teachers:")
			for _, t := range data.Teachers {
				fmt.Printf("Name: %s\nSubject: %s\nPosition: %s\nHome City: %s\n", t.Name, t.Subject, t.position, t.HomeCity)
			}
			fmt.Println("Staff:")
			for _, st := range data.Staff {
				fmt.Printf("Name: %s\nMax Salary: %d\nHome City: %s\n", st.Name, st.MaxSalary, st.HomeCity)
			}
			fmt.Println("________________________")

		case 5:
			fmt.Println("Search College Details:")
			fmt.Println("1 - Search Students by Name")
			fmt.Println("2 - Search Teachers by Subject")
			fmt.Println("3 - Search Staff by City")
			var searchChoice int
			fmt.Scan(&searchChoice)
			readLine()

			switch searchChoice {
			case 1:
				fmt.Println("Enter student name to search:")
				var name string
				name = readLine()
				fmt.Println("Matching Students:")
				found := false
				for _, s := range data.Students {
					if s.Name == name {
						fmt.Printf("Name: %s\nCourse: %s\nBranch: %s\nAge: %d\nCGPA: %f\n", s.Name, s.Course, s.Branch, s.Age, s.CGPA)
						found = true
					}
				}
				if !found {
					fmt.Println("No matching student found.")
				}

			case 2:
				fmt.Println("Enter teacher subject to search:")
				var subject string
				subject = readLine()
				fmt.Println("Matching Teachers:")
				found := false
				for _, t := range data.Teachers {
					if t.Subject == subject {
						fmt.Printf("Name: %s\nSubject: %s\nPosition: %s\nHome City: %s\n", t.Name, t.Subject, t.position, t.HomeCity)
						found = true
					}
				}
				if !found {
					fmt.Println("No matching teacher found.")
				}

			case 3:
				fmt.Println("Enter staff city to search:")
				var city string
				fmt.Scan(&city)
				fmt.Println("Matching Staff:")
				for _, st := range data.Staff {
					if st.HomeCity == city {
						fmt.Printf("Name: %s\nMax Salary: %d\nHome City: %s\n", st.Name, st.MaxSalary, st.HomeCity)
					}
				}

			default:
				fmt.Println("Invalid search option.")
			}
		case 6:
			fmt.Println("Thank you for using the College Detail System. Have a great day!")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
