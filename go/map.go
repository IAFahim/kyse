package main

import (
	"fmt"
)

func main() {
	_ = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	_ = []int{1}
	var s string
	s = "{alt}+abc{ax}+{win}+{alt}"
	arr := make([]int, len(s)/2+1)
	i := 0
	if len(s) > 0 && s[0] == '{' {
		i += specialKey(s, 0)
	}
	for ; i < len(s); i++ {
		fmt.Println(string(s[i]))
		if s[i] == '+' {
			i += specialKey(s, i+1)
		}

	}
	fmt.Printf("%v", arr)
}

func specialKey(s string, i int) (offset int) {
	offset = 0
	if s[i] == '{' {
		x := i + 1
		for ; x < len(s) && s[x] != '}'; x++ {
		}
		switch x - i - 1 {
		case 3:
			switch s[i+1] {
			case 'a', 'A':
				fmt.Println("hasALT")
			case 'w', 'W':
				fmt.Println("hasWin")
			}
			offset += 3
		case 5:
			switch s[i+1] {
			case 's':
				if s[i+2] == 'h' {
					fmt.Println("hasShift")
				} else if s[i+2] == 'u' {
					fmt.Println("hasSuper")
				}
			case 'a':
				fmt.Println("hasAltGr")
			}
		}
		offset += 2
	}
	return offset
}
