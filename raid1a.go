package piscine

import "fmt"


func printline( n,x,y int) {
	if(x == 1 || x==y){
		fmt.Print("o")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print("-")
		}
		if (n >=2 || n < 1){
			fmt.Print("o")
		}
		fmt.Print("\n")
	} else {
		fmt.Print("|")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print(" ")
		}
		if n >=2 {
			fmt.Print("|")
		}
		fmt.Print("\n")
	}
}

func Raid1a(x,y int)  {
	for i:=1;i<=y;i++{
		printline(x,i,y)
	}
}
