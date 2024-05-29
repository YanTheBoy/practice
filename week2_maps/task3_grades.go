package week2_maps

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade,omitempty"`
}

type Object struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	ObjectID  int `json:"object_id"`
	StudentID int `json:"student_id"`
	Result    int `json:"result"`
}

type Register struct {
	Students []Student
	Objects  []Object
	Results  []Result
}

func ParseDzFile(s map[int]Student, o map[int]Object, r Register) (map[int]Student, map[int]Object, Register) {

	file, err := os.ReadFile("dz3.json")
	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}
	if err = json.Unmarshal(file, &r); err != nil {
		log.Fatalf("Cannot unmarshal data: %v", err)
	}

	for _, student := range r.Students {
		s[student.Id] = student
	}
	for _, object := range r.Objects {
		o[object.Id] = object
	}

	return s, o, r
}

func CalcResult() {
	var register Register
	var students = make(map[int]Student)
	var objects = make(map[int]Object)

	students, objects, register = ParseDzFile(students, objects, register)

	w := tabwriter.NewWriter(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.TabIndent)

	fmt.Fprintln(w, "___________________________________________\nStudent Name\t   Grade\t  Object\t  Result\n___________________________________________")

	for _, result := range register.Results {
		student := students[result.StudentID]
		object := objects[result.ObjectID]

		fmt.Fprintf(w, "%s\t    %d\t  %s\t  %d\n", student.Name, student.Grade, object.Name, result.Result)

	}
	w.Flush()

}

/* 3. У учеников старших классов прошел контрольный срез по нескольким предметам. Выведите данные в читаемом виде
в таблицу вида
_____________________________________
Student name  | Grade | Object    |   Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
...
Вводные данные представлены в файле dz3.json
*/
