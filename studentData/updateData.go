package studentData

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

func UpdateDataValue(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	job := r.FormValue("job")
	roll := r.FormValue("rollno")
	skills := r.FormValue("skills")

	type Student struct {
		Name   string   `yaml: "name"`
		Roll   string   `yaml: "roll"`
		Job    string   `yaml: "job"`
		Skills []string `yaml: "skills"`
	}

	studentsData := &[]Student{}

	source, err := ioutil.ReadFile("../student_data.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData)
	if err != nil {
		log.Println(err)
	}

	rollno, err := strconv.Atoi(roll)
	rollno--
	if name != "" {
		(*studentsData)[rollno].Name = name
	}
	if job != "" {
		(*studentsData)[rollno].Job = job
	}
	if skills != "" {
		skillsarray := strings.Split(skills, ",")
		fmt.Println(skillsarray)
		for _, skill := range skillsarray {

			(*studentsData)[rollno].Skills = append((*studentsData)[rollno].Skills, skill)

		}

	}

	d, err := yaml.Marshal(&studentsData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile("../student_data.yaml", d, 0644)

	http.Redirect(w, r, "/", http.StatusSeeOther)

	uri := r.URL.String()
	method := r.Method
	fmt.Println(uri, method)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	type Student struct {
		Name   string   `yaml: "name"`
		Roll   string   `yaml: "roll"`
		Job    string   `yaml: "job"`
		Skills []string `yaml: "skills"`
	}

	studentsData := &[]Student{}

	source, err := ioutil.ReadFile("../student_data.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData)
	if err != nil {
		log.Println(err)
	}

	Rollno := r.FormValue("first")
	rollno, err := strconv.Atoi(Rollno)

	if len(*studentsData) < rollno || (Rollno == "") {
		http.Redirect(w, r, "/", http.StatusSeeOther)

		log.Println(err)
	} else {
		type rolldata struct {
			RollNo string
		}

		rolldata1 := rolldata{RollNo: Rollno}

		t, _ := template.ParseFiles("../update.html")
		t.Execute(w, rolldata1)
	}

	uri := r.URL.String()
	method := r.Method
	fmt.Println(uri, method)

}
