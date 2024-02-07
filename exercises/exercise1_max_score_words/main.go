package main

import (
	"fmt"
	"slices"
	"sync"
)

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func generateCombinations(elements []string) [][]string {
	var result [][]string
	generateHelper(elements, 0, &result)
	return result
}

func generateHelper(elements []string, start int, result *[][]string) {
	if start == len(elements)-1 {
		combination := make([]string, len(elements))
		copy(combination, elements)
		*result = append(*result, combination)
		return
	}

	for i := start; i < len(elements); i++ {
		elements[start], elements[i] = elements[i], elements[start]
		generateHelper(elements, start+1, result)
		elements[start], elements[i] = elements[i], elements[start]
	}
}

func indexOf(element byte, data []byte) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

var wg sync.WaitGroup

func maxScoreWords(words []string, letters []byte, score []int) int {
	var alphabetMap map[byte]int = make(map[byte]int)
	var alphabet []byte = []byte("abcdefghijklmnopqrstuvwxyz")
	var wordsScoredMap map[string]int = make(map[string]int)

	for i := 0; i < len(alphabet); i += 1 {
		alphabetMap[alphabet[i]] = score[i]
	}

	var wordsTemp = removeDuplicateStr(words)
	for _, word := range wordsTemp {
		for _, l := range []byte(word) {
			if _, v := wordsScoredMap[word]; !v {
				wordsScoredMap[word] = alphabetMap[l]
			} else {
				wordsScoredMap[word] += alphabetMap[l]
			}
		}

	}
	sumsChan := make(chan int)
	wordsCombinations := generateCombinations(words)
	for _, words := range wordsCombinations {
		go func(words []string) {
			wg.Add(1)

			var results []string = []string{""}
			var f bool = true
			var scoreSum int = 0
			var _letters []byte
			_letters = append(_letters, letters...)

			for _, word := range words {
				for _, l := range []byte(word) {
					if slices.Contains(_letters, l) {
						var index = indexOf(l, _letters)
						_letters = append(_letters[:index], _letters[index+1:]...)
					} else {
						f = false
						break
					}
				}
				if f {
					results = append(results, word)
				}

			}
			for _, w := range results {
				scoreSum += wordsScoredMap[w]
			}
			sumsChan <- scoreSum
			defer wg.Done()
		}(words)
	}
	var scoredSums []int
	for i := 0; i < len(wordsCombinations); i += 1 {
		s := <-sumsChan
		scoredSums = append(scoredSums, s)
	}
	wg.Wait()
	close(sumsChan)
	return slices.Max(scoredSums)
}

func maxScoreWords_one_thread(words []string, letters []byte, score []int) int {
	var alphabetMap map[byte]int = make(map[byte]int)
	var alphabet []byte = []byte("abcdefghijklmnopqrstuvwxyz")
	var wordsScoredMap map[string]int = make(map[string]int)

	for i := 0; i < len(alphabet); i += 1 {
		alphabetMap[alphabet[i]] = score[i]
	}

	var wordsTemp = removeDuplicateStr(words)
	for _, word := range wordsTemp {
		for _, l := range []byte(word) {
			if _, v := wordsScoredMap[word]; !v {
				wordsScoredMap[word] = alphabetMap[l]
			} else {
				wordsScoredMap[word] += alphabetMap[l]
			}
		}

	}
	var scoredSums []int
	wordsCombinations := generateCombinations(words)
	for _, words := range wordsCombinations {

		var results []string = []string{""}
		var f bool = true
		var scoreSum int = 0
		var _letters []byte
		_letters = append(_letters, letters...)

		for _, word := range words {
			for _, l := range []byte(word) {
				if slices.Contains(_letters, l) {
					var index = indexOf(l, _letters)
					_letters = append(_letters[:index], _letters[index+1:]...)
				} else {
					f = false
					break
				}
			}
			if f {
				results = append(results, word)
			}

		}
		for _, w := range results {
			scoreSum += wordsScoredMap[w]
		}
		scoredSums = append(scoredSums, scoreSum)
	}

	return slices.Max(scoredSums)
}

func main() {
	w := []string{"dog", "cat", "dad", "good"}
	l := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	sc := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(maxScoreWords(w, l, sc))

	w = []string{"xxxz", "ax", "bx", "cx"}
	l = []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	sc = []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(maxScoreWords(w, l, sc))

	w = []string{"ad", "dbacbbedc", "ae", "adbdacad", "dcdecacdcb", "ddbba", "dbcdbeaade", "aeccdcb", "bce"}
	l = []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'd', 'd', 'd', 'd', 'e', 'e', 'e', 'e', 'e', 'e'}
	sc = []int{1, 8, 3, 1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(maxScoreWords(w, l, sc))

	w = []string{"dog", "cat", "dad", "good"}
	l = []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	sc = []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(maxScoreWords_one_thread(w, l, sc))

	w = []string{"xxxz", "ax", "bx", "cx"}
	l = []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	sc = []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(maxScoreWords_one_thread(w, l, sc))

	w = []string{"ad", "dbacbbedc", "ae", "adbdacad", "dcdecacdcb", "ddbba", "dbcdbeaade", "aeccdcb", "bce"}
	l = []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'd', 'd', 'd', 'd', 'e', 'e', 'e', 'e', 'e', 'e'}
	sc = []int{1, 8, 3, 1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(maxScoreWords_one_thread(w, l, sc))

}

// TODO one thread is to slow
// TODO multi thread use to much memory
