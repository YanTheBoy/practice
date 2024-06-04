package week2_maps

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type Cache[K comparable, V any] struct {
	m map[K]V
}

func (c *Cache[K, V]) Init() {
	c.m = make(map[K]V)
}
func (c *Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}
func (c *Cache[K, V]) Get(key K) (V, bool) {
	k, ok := c.m[key]
	return k, ok
}

func ParseDzFileWithCache() (Register, Cache[int, Student], Cache[int, Object]) {
	var r Register
	var students Cache[int, Student]
	var objects Cache[int, Object]
	students.Init()
	objects.Init()

	file, err := os.ReadFile("dz3.json")
	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}
	if err = json.Unmarshal(file, &r); err != nil {
		log.Fatalf("Cannot unmarshal data: %v", err)
	}

	for _, student := range r.Students {
		students.Set(student.Id, student)
	}
	for _, object := range r.Objects {
		objects.Set(object.Id, object)
	}
	return r, students, objects

}

func CalcResult2() {
	register, students, objects := ParseDzFileWithCache()

	w := tabwriter.NewWriter(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.TabIndent)

	fmt.Fprintln(w, "___________________________________________\nStudent Name\t   Grade\t  Object\t  Result\n___________________________________________")

	for _, result := range register.Results {

		student, _ := students.Get(result.StudentID)
		object, _ := objects.Get(result.ObjectID)

		fmt.Fprintf(w, "%s\t    %d\t  %s\t  %d\n", student.Name, student.Grade, object.Name, result.Result)

	}
	w.Flush()

}

/* 5. Перепишите задачу #3 с использованием структуры-дженерик Cache, изученной на семинаре.
Храните в кеше таблицы студентов и предметов.
*/
