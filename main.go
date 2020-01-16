package main

import (
	"encoding/json"
	"fmt"

	"gitlab.com/shubhashree/practise/student/config"
	"gitlab.com/shubhashree/practise/student/datagen"
)

func main() {
	//fmt.Println("Inside main")
	s := make([]*config.Student, 0)

	//SampleTest()
	for count := 0; count < 2000; count++ {
		s = append(s, datagen.GetFakeStudentWithAllSem())
	}

	jsonData, _ := json.Marshal(s)
	fmt.Println(string(jsonData))
}

func factoryPattern() {
	s1 := config.FactoryMethod("Aaron", "001").
		AddSemester("1st Semester").
		AddSubject("1st Semester", "AI", 99).
		AddSubject("1st Semester", "os", 99).
		AddSemester("2nd Semester").
		AddSubject("2nd Semester", "AI", 99).
		AddSubject("2nd Semester", "os", 99)

	jsonData, _ := json.Marshal(s1)
	fmt.Println(string(jsonData))
}

func SampleTest() {
	s1 := config.Student{}
	s1.Name = "Aaron"
	s1.Rollno = "001"

	subject1 := make([]*config.Subject, 1)
	//subject1 = append(subject1, config.Subject{})
	subject1[0] = &config.Subject{}
	subject1[0].Name = "AI"
	subject1[0].Marks = 99

	s1.Semester = make(map[string]*config.Courses)
	s1.Semester["1st_sem"] = &config.Courses{
		Semester: "first",
		Subjects: subject1,
	}
	fmt.Println("Student info :", s1)

	jsonData, _ := json.Marshal(s1)
	fmt.Println(string(jsonData))
}

func sampleTest2() {
	s1 := config.FactoryMethod("Aaron", "001")
	fmt.Println("Student info :", s1)
	s1.Semester["1st Semester"] = &config.Courses{}
	s1.Semester["1st Semester"].Semester = "1st Semester"
	s1.Semester["1st Semester"].Subjects = make([]*config.Subject, 2)
	s1.Semester["1st Semester"].Subjects[0] = &config.Subject{}
	s1.Semester["1st Semester"].Subjects[0].Name = "AI"
	s1.Semester["1st Semester"].Subjects[0].Marks = 99
	s1.Semester["1st Semester"].Subjects[1] = &config.Subject{}
	s1.Semester["1st Semester"].Subjects[1].Name = "OS"
	s1.Semester["1st Semester"].Subjects[1].Marks = 98
	fmt.Println("Student info :", s1)

	jsonData, _ := json.Marshal(s1)
	fmt.Println(string(jsonData))
}
