package main

// Model  for course  -- file

type Course struct {
	CourseId string `json:"coursed"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db

var courses []Course;

// middleware , helper - file 

func (c * Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
