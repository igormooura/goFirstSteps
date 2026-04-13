package main

import (
	"fmt"
	"sort"
)

type Activity struct {
	start int
	end   int
}

func activitySelection(activity []Activity) []Activity {

	sort.Slice(activity, func(i, j int) bool { // O(n log n)
		return activity[i].end < activity[j].end
	})

	var result []Activity

	// primeira atividadade selecionada
	result = append(result, activity[0]) // O(1)
	lastEnd := activity[0].end

	// Percorre o resto
	for i := 1; i < len(activity); i++ { // O(n)
		if activity[i].start >= lastEnd { // O(1)
			result = append(result, activity[i]) // O(1)
			lastEnd = activity[i].end            // O(1)
		}
	}

	return result
}

func main() {
	activities := []Activity{
		{1, 3},
		{2, 5},
		{4, 7},
		{6, 9},
		{8, 10},
		{9, 11},
	}

	selected := activitySelection(activities)

	fmt.Println("Atividades selecionadas:")
	for _, a := range selected {
		fmt.Printf("(%d, %d)\n", a.start, a.end)
	}
}
