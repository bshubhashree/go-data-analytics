package datagen

import (
	"fmt"

	"gitlab.com/shubhashree/practise/student/config"
	"syreclabs.com/go/faker"
)

var (
	semister = []string{"1st_sem", "2nd_sem", "3rd_sem", "4th_sem"}
	subjects = map[string][]string{
		"1st_sem": []string{"English", "Programming", "physics", "Chemestry"},
		"2nd_sem": []string{"Discrete Mathematics", "Data Structures", "Computer Architecture", "Database Management"},
		"3rd_sem": []string{"Computer Networks", "Analysis of Algorithms", "Operating Systems", "Software Engineering"},
		"4th_sem": []string{"Computer Graphics", "Data Analysis", "Web Application Development", "Software Quality"},
	}
)

func GetFakeStudent() *config.Student {
	fakePerson := faker.Name()
	s1 := config.FactoryMethod(fakePerson.Name(), fmt.Sprintf("A20FT%s", faker.Number().Number(5)))

	s1.FirstName = fakePerson.FirstName()
	s1.LastName = fakePerson.LastName()
	s1.Prefix = fakePerson.Prefix()

	//s1.About = faker.Hacker().Phrases()
	// jsonData, _ := json.Marshal(s1)
	// fmt.Println("\n\n\n", string(jsonData))
	return s1
}

func GetFakeStudentWithAllSem() *config.Student {
	fakePerson := faker.Name()
	s1 := config.FactoryMethod(fakePerson.Name(), fmt.Sprintf("A20FT%s", faker.Number().Number(5)))

	s1.FirstName = fakePerson.FirstName()
	s1.LastName = fakePerson.LastName()
	s1.Prefix = fakePerson.Prefix()

	//s1.About = faker.Hacker().Phrases()
	// jsonData, _ := json.Marshal(s1)
	// fmt.Println("\n\n\n", string(jsonData))
	for _, sem := range semister {
		s1.AddSemester(sem)
		for _, subject := range subjects[sem] {
			s1.AddSubject(sem, subject, faker.Number().NumberInt(2))
		}
	}
	return s1
}
