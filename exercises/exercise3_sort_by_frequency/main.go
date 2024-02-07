package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string {
	var fMap = make(map[string]int)

	for _, ch := range strings.Split(s, "") {
		if _, value := fMap[ch]; value {
			fMap[ch] += 1
		} else {
			fMap[ch] = 1
		}
	}

	keys := make([]string, 0, len(fMap))
	for key := range fMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return fMap[keys[i]] > fMap[keys[j]]
	})

	var results = ""
	for _, k := range keys {
		for i := 0; i < fMap[k]; i++ {
			results += k
		}
	}
	return results

}

func main() {
	fmt.Println(frequencySort("tree"))   //eert
	fmt.Println(frequencySort("cccaaa")) //aaaccc or cccaaa
	fmt.Println(frequencySort("Aabb"))   //bbAa

}
