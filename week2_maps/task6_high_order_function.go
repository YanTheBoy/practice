package week2_maps

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"
)

func CalculateMeanResults() {
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

		var meanValuesPerGrade []float32

		fmt.Fprintf(w, "_________________\n%s\t Mean\n_________________\n", object.Name)

		for grade, studentCount := range grades {

			fmt.Fprintf(w, "%d grade\t", grade)
			var meanValue float32
			var meanValues []float32

			filteredReg := Filter(register.Results, func(r Result) bool {
				if students[r.StudentID].Grade == grade && r.ObjectID == object.Id {
					return true
				}
				return false
			})

			// Заполняем срез
			for _, res := range filteredReg {
				meanValues = append(meanValues, float32(res.Result))
			}

			meanValue = Reduce(meanValues, 0, func(a, b float32) float32 {
				return a + b
			})

			meanValuePerGrade := meanValue / float32(studentCount)
			meanValuesPerGrade = append(meanValuesPerGrade, meanValuePerGrade)

			fmt.Fprintf(w, " %.1f\n", meanValuePerGrade)
		}

		meanValuePerObject = Reduce(meanValuesPerGrade, 0, func(a, b float32) float32 {
			return a + b
		})

		fmt.Fprintf(w, "_________________\nmean\t %.1f\n_________________\n", meanValuePerObject/float32(len(grades)))

	}

	w.Flush()
}

func Reduce[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	r := init
	for _, v := range s {
		r = f(v, r)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

/* 6. Перепишите задачу #4 с использованием функций высшего порядка, изученных на лекции.
Желательно реализуйте эти функции самостоятельно.
*/
