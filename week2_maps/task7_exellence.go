package week2_maps

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func ShowOnlyExcellentStudents() {
	register, students, objects := ParseDzFileWithCache()

	onlyExcellentResults := Filter(register.Results, func(r Result) bool {
		return r.Result == 5
	})

	w := tabwriter.NewWriter(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.TabIndent)

	fmt.Fprintln(w, "___________________________________________\nStudent Name\t   Grade\t  Object\t  Result\n___________________________________________")

	for _, result := range onlyExcellentResults {
		student, _ := students.Get(result.StudentID)
		object, _ := objects.Get(result.ObjectID)

		fmt.Fprintf(w, "%s\t    %d\t  %s\t  %d\n", student.Name, student.Grade, object.Name, result.Result)

	}
	w.Flush()

}

/*
7. Выведите в консоль круглых отличников из числа студентов, используя функцию Filter.
Вывод реализуйте как в задаче #3.
_____________________________________
Student name  | Grade | Object    |   Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
*/
