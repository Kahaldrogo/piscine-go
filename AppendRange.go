package piscine_go

import (
	"fmt"
	"strings"
)

func AppendRange(min,max int) []int {
	var aux [] int
	for i:=min; i<max; i++{
		aux = append(aux, i)
	}
	return aux
}

func MakeRange(min,max int) []int {
	if min>=max{
		var aux [] int
		return aux
	} else {
		aux := make([]int,max-min)
		for i:=0; i<len(aux); i++{
			aux[i] = min+i
		}
		return aux
	}
}

func ConcatParams(args []string) string {
	return strings.Join(args, "\n")
}

func SplitWhiteSpaces(str string) []string {
	return strings.Fields(str)
}

func PrintWordsTables(table []string) {
	for _, mot := range table{
		fmt.Println(mot)
	}
}