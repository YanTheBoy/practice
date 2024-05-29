package week2_maps

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"
)

func CalcMeanResults() {
	var register Register
	var students = make(map[int]Student)
	var objects = make(map[int]Object)

	students, objects, register = ParseDzFile(students, objects, register)

	w := tabwriter.NewWriter(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.TabIndent)

	grades := make(map[int]int)
	for _, stuData := range register.Students {
		grades[stuData.Grade]++
	}

	// Сотрируем мапу по ключам
	keys := make([]int, 0, len(grades))
	for k := range grades {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, object := range objects {
		var meanValuePerObject float32

		fmt.Fprintf(w, "_________________\n%s\t Mean\n_________________\n", object.Name)

		for grade, studentCount := range grades {

			fmt.Fprintf(w, "%d grade\t", grade)
			var meanValue float32
			for _, res := range register.Results {
				if students[res.StudentID].Grade == grade && res.ObjectID == object.Id {
					meanValue += float32(res.Result)
				}
			}

			meanValuePerGrade := meanValue / float32(studentCount)
			meanValuePerObject += meanValuePerGrade
			fmt.Fprintf(w, " %.1f\n", meanValuePerGrade)
		}
		fmt.Fprintf(w, "_________________\nmean\t %.1f\n_________________\n", meanValuePerObject/float32(len(grades)))

	}

	w.Flush()
}

/* 4. Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам в виде:
________________
Math	 | Mean
________________
 9 grade | 4.5
10 grade | 5
11 grade | 3.5
________________
mean     | 4		- среднее значение среди всех учеников
________________
________________
Biology	 | Mean
________________
...
Вводные данные представлены в файле dz3.json
*/
