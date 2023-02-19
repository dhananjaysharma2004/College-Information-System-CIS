package main

import (
	"fmt"
	"sync"
	"time"
)

// Students, Teachers and Staff Details Project
type studentdetails struct {
	name   string
	course string
	branch string
	age    int
	cgpa   float32
}
type teacherdetails struct {
	name      string
	subject   string
	maxsalary int
	homecity  string
}
type staffdetails struct {
	name      string
	maxsalary int
	homecity  string
}

var waitgroup = sync.WaitGroup{}

func main() {
	var n int
	for {
		fmt.Println("WELCOME LADIES AND GENTLEMENT!!!")
		fmt.Println("Please enter 1-students, 2-teachers & 3-staff")
		fmt.Scan(&n)

		if n == 1 {
			var stu studentdetails
			fmt.Println("Enter NAME of the student")
			fmt.Scan(&stu.name)
			fmt.Println("Enter COURSE of the student")
			fmt.Scan(&stu.course)
			fmt.Println("Enter BRANCH of the student")
			fmt.Scan(&stu.branch)
			fmt.Println("Enter AGE of the student")
			fmt.Scan(&stu.age)
			fmt.Println("Enter CGPA of the student")
			fmt.Scan(&stu.cgpa)
			waitgroup.Add(2)
			var stuname string = stu.name
			var stucourse string = stu.course
			var stubranch string = stu.branch
			var stuage int = stu.age
			var stucgpa float32 = stu.cgpa
			go printdetails1(stuname, stucourse, stubranch, stuage, stucgpa)
		} else if n == 2 {
			var tea teacherdetails
			fmt.Println("Enter NAME of the teacher")
			fmt.Scan(&tea.name)
			fmt.Println("Enter SUBJECT of the teacher")
			fmt.Scan(&tea.subject)
			fmt.Println("Enter MAXSALARY of the teacher")
			fmt.Scan(&tea.maxsalary)
			fmt.Println("Enter HOMECITY of the teacher")
			fmt.Scan(&tea.homecity)
			waitgroup.Add(2)
			var teaname string = tea.name
			var teasubject string = tea.subject
			var teamaxsalary int = tea.maxsalary
			var teahomecity string = tea.homecity
			go printdetails2(teaname, teasubject, teamaxsalary, teahomecity)
		} else if n == 3 {
			var sta staffdetails
			fmt.Println("Enter NAME of the staff")
			fmt.Scan(&sta.name)
			fmt.Println("Enter MAXSALARY of the staff")
			fmt.Scan(&sta.maxsalary)
			fmt.Println("Enter HOMECITY of the staff")
			fmt.Scan(&sta.homecity)
			waitgroup.Add(2)
			var staffname string = sta.name
			var staffmaxsalary int = sta.maxsalary
			var staffhomecity string = sta.homecity
			go printdetails3(staffname, staffmaxsalary, staffhomecity)
		}
	}
}
func printdetails1(stuname string, stucourse string, stubranch string, stuage int, stucgpa float32) {
	time.Sleep(30 * time.Second)
	fmt.Println("=========================================================================================")
	fmt.Printf("NAME OF THE STUDENT : %v\nCOURSE : %v\nBRANCH : %v\nAGE : %v\nCGPA : %v\n", stuname, stucourse, stubranch, stuage, stucgpa)
	fmt.Println("THANK YOU FOR VISITING OUR SITE!!! HAVE A GOOD DAY AHEAD")
	fmt.Println("=========================================================================================")
	waitgroup.Done()
}

func printdetails2(teaname string, teasubject string, teamaxsalary int, teahomecity string) {
	time.Sleep(30 * time.Second)
	fmt.Println("=========================================================================================")
	fmt.Printf("NAME OF THE TEACHER : %v\nSUBJECT : %v\nMAXSALARY : %v\nHOMECITY : %v\n", teaname, teasubject, teamaxsalary, teahomecity)
	fmt.Println("THANK YOU FOR VISITING OUR SITE!!! HAVE A GOOD DAY AHEAD")
	fmt.Println("=========================================================================================")
	waitgroup.Done()
}

func printdetails3(staffname string, staffmaxsalary int, staffhomecity string) {
	time.Sleep(30 * time.Second)
	fmt.Println("=========================================================================================")
	fmt.Printf("NAME OF THE STAFF : %v\nMAXSALARY : %v\nHOMECITY : %v\n", staffname, staffmaxsalary, staffhomecity)
	fmt.Println("THANK YOU FOR VISITING OUR SITE!!! HAVE A GOOD DAY AHEAD")
	fmt.Println("=========================================================================================")
	waitgroup.Done()
}
