package main

import (
	"fmt"
	// "sort"
)

type pair struct{
	name string
	count int
}

func mostUpperItem(items []string) []pair {
	counts := make(map[string]int)
	for _, str := range items {
		counts[str]++
	}

	var results []pair
	for str, cnt := range counts {
		results = append(results, pair{name: str, count: cnt})
	}
	// implementasi package sort
	// sort.Slice(results, func(i, j int) bool {
	// 	return results[i].count > results[j].count
	// })

	n := len(results)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if results[j].count > results[j+1].count {
				results[j], results[j+1] = results[j+1], results[j]
			}
		}
	}

	//output sesuai instruksi namun return value harus berupa string
	// var output string
	// for _, p := range results {
	// 	output += fmt.Sprintf("%s->%d, ", p.name, p.count)
	// }

	// return output[:len(output)-2]	

	return results
}

func main()  {
	fmt.Println(mostUpperItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(mostUpperItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(mostUpperItem([]string{"football", "basketball", "tennis"}))
}