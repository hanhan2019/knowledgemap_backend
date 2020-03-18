package model

func (people *People) GetAllInform() {

}

func (people *People) CheackPassWord(_passWord string) bool {
	if people.Password != _passWord {
		return false
	}
	return true
}

func (people *People) ChangePassWord(_passWord string) {
	people.Password = _passWord
}

func (student *Student) ChangeOrigin(_origin string) {
	student.Origin = _origin
}

func (student *Student) ChangeClass(_class string) {
	student.Class = _class
}

func (student *Student) AddCourse(_courses []string) {
	student.Courses = append(student.Courses, _courses...)
}

func (student *Student) DeleteCourse(_course string) bool {
	var j int = -1
	for i, v := range student.Courses {
		if v == _course {
			j = i
		}
	}
	if j == -1 {
		return false
	} else {
		student.Courses = append(student.Courses[:j], student.Courses[j+1:]...)
	}
	return true
}

func (student *Student) ChangeCourse(_course string) bool {
	var j int = -1
	for i, v := range student.Courses {
		if v == _course {
			j = i
		}
	}
	if j == -1 {
		return false
	} else {
		student.Courses[j] = _course
	}
	return true
}

// func (class *Class)GetAllStudent()
