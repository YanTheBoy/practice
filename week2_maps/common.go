package week2_maps

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


