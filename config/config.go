package config

type Student struct {
	Name      string
	FirstName string
	LastName  string
	Prefix    string
	About     []string
	Rollno    string
	Semester  map[string]*Courses
}

type Courses struct {
	Semester string
	Subjects []*Subject
}

type Subject struct {
	Name  string
	Marks int
}

func FactoryMethod(name, rollNo string) *Student {
	s1 := Student{}
	s1.Name = name
	s1.Rollno = rollNo
	// subject1 := make([]*Subject, 0)
	s1.Semester = make(map[string]*Courses)
	// s1.Semester["1st_sem"] = &Courses{}
	return &s1
}

func (s *Student) AddSemester(name string) *Student {
	s.Semester[name] = &Courses{
		Semester: name,
	}
	return s
}

func (s *Student) AddSubject(semesterName, name string, marks int) *Student {

	if len(s.Semester[semesterName].Subjects) == 0 {
		s.Semester[semesterName].Subjects = make([]*Subject, 0)
	}
	s.Semester[semesterName].Subjects = append(s.Semester[semesterName].Subjects, &Subject{})
	// s.Semester[semesterName].Subjects[0] = &Subject{}
	s.Semester[semesterName].Subjects[len(s.Semester[semesterName].Subjects)-1].Name = name
	s.Semester[semesterName].Subjects[len(s.Semester[semesterName].Subjects)-1].Marks = marks
	return s
}
