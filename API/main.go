package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Model  for course  -- file

type Course struct {
	CourseId    string `json:"coursed"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author      *Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db

var courses []Course

// middleware , helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API CALLING")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 299, Author: &Author{
		Fullname: "Saksam chandel",
		Website:  "chandel.social",
	}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 199, Author: &Author{
		Fullname: "jack",
		Website:  "reddit",
	}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getALLCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeletOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":3000", r))
}

// controllers - file

// serve home  route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Buliding FIRST API USING GO</h1>"))
}

func getALLCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req
	params := mux.Vars(r)

	// find matching id and retur it

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// empty body

	if r.Body == nil {
		json.NewEncoder(w).Encode("No Data")
	}

	// {}

	var NewCourse Course
	_ = json.NewDecoder(r.Body).Decode(&NewCourse)

	if NewCourse.IsEmpty() {
		json.NewEncoder(w).Encode("No Data inside it")
		return
	}
	for _, course := range courses {
		if course.CourseName == NewCourse.CourseName {
			json.NewEncoder(w).Encode("Course name already exists")
			return
		}
	}
	// generate unique id , string
	rand.Seed(time.Now().UnixNano())
	NewCourse.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, NewCourse)
	json.NewEncoder(w).Encode(NewCourse)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func DeletOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delet one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
