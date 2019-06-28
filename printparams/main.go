package main

import (
	"fmt"
	"os"
)

func main()  {
	aux := os.Args
	for i:=1; i<len(aux);i++{
		fmt.Println(aux[i])
	}
}