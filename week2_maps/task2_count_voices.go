package week2_maps

import (
	"sort"
)

/* 2. Подсчет голосов.
Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
Пример.
Вход: ["Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"]
Вывод: [{Ann, 3}, {Kate, 2}, {Peter, 1}, {Helen, 1}]
*/

type Candidate struct {
	Name  string
	Votes int
}

func CountVoices(names []string) []Candidate {
	countedVoices := make(map[string]int)
	for _, name := range names {
		countedVoices[name]++
	}

	var candidates []Candidate

	for name, voices := range countedVoices {
		candidates = append(candidates, Candidate{
			Name:  name,
			Votes: voices,
		})
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Votes > candidates[j].Votes
	})

	return candidates
}
