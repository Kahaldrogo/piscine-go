package main

import (
	"fmt"
	"os"
)

func main()  {
	aux := os.Args
	for i:=len(aux)-1; i>0;i--{
		fmt.Println(aux[i])
	}
}