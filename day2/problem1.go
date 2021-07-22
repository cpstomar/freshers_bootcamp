package main
import (
	"fmt"
	"strings"
)
func Count(str string, results chan<- [26]int) {
	var cnt [26]int
	for i := 0; i < 26; i++ {
		cnt[i] += strings.Count(str, string(rune('a'+i)))
	}
	results <- cnt
}
func main() {
	numberOfStrings := 5
	Input := make([]string, 0)
	Results := make(chan [26]int, numberOfStrings)
	Input = append(Input, "quick")
	Input = append(Input, "brown")
	Input = append(Input, "fox")
	Input = append(Input, "lazy")
	Input = append(Input, "add")

	for i := 0; i < 5; i++ {
		go Count(Input[i], Results)
	}
	var finalFrequency [26]int
	for i := 0; i < numberOfStrings; i++ {
		result := <-Results
		for i := 0; i < 26; i++ {
			finalFrequency[i] += result[i]
		}
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", rune('a'+i), finalFrequency[i])
	}
}