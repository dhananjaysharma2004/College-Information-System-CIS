# College Detail System

The College Detail System is a Go program that manages and stores information about students, teachers, and staff in a university. The data is stored in a JSON file named university.json.

## Features

- Add student details including name, course, branch, age, and CGPA.
- Add teacher details including name, subject, position, and home city.
- Add staff details including name, max salary, and home city.
- View university details, including lists of students, teachers, and staff.
- Search for specific details such as students by name, teachers by subject, and staff by city.

## Usage

1. _Add Student:_

   - Select option 1 and provide the required details for the student.

2. _Add Teacher:_

   - Select option 2 and provide the required details for the teacher.

3. _Add Staff:_

   - Select option 3 and provide the required details for the staff.

4. _View Details:_

   - Select option 4 to view the details of all students, teachers, and staff.

5. _Search Details:_

   - Select option 5 to search for specific details based on the type (students, teachers, staff).
     - For students, you can search by name.
     - For teachers, you can search by subject.
     - For staff, you can search by city.

6. _Exit:_
   - Select option 6 to exit the program.

## How to Run

1. Ensure you have Go installed on your machine.
2. Clone the repository.
3. Run the main.go file using the command: go run main.go

## Data Storage

The program uses a JSON file named university.json to persistently store the university's data.

## License

This project is licensed under the [MIT License](LICENSE).

---

Thank you for using the College Detail System!
