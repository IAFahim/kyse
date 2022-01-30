package main

import (
	"fmt"
	"github.com/micmonay/keybd_event"
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
		if s[i] == '+' {
			i += specialKey(s, i+1)
		}
	}
	fmt.Printf("%v", arr)
}

type keySet struct {
	code  int
	shift bool
}

var (
	names = map[string]keySet{
		"a": {keybd_event.VK_A, false},
		"b": {keybd_event.VK_B, false},
		"c": {keybd_event.VK_C, false},
		"d": {keybd_event.VK_D, false},
		"e": {keybd_event.VK_E, false},
		"f": {keybd_event.VK_F, false},
		"g": {keybd_event.VK_G, false},
		"h": {keybd_event.VK_H, false},
		"i": {keybd_event.VK_I, false},
		"j": {keybd_event.VK_J, false},
		"k": {keybd_event.VK_K, false},
		"l": {keybd_event.VK_L, false},
		"m": {keybd_event.VK_M, false},
		"n": {keybd_event.VK_N, false},
		"o": {keybd_event.VK_O, false},
		"p": {keybd_event.VK_P, false},
		"q": {keybd_event.VK_Q, false},
		"r": {keybd_event.VK_R, false},
		"s": {keybd_event.VK_S, false},
		"t": {keybd_event.VK_T, false},
		"u": {keybd_event.VK_U, false},
		"v": {keybd_event.VK_V, false},
		"w": {keybd_event.VK_W, false},
		"x": {keybd_event.VK_X, false},
		"y": {keybd_event.VK_Y, false},
		"z": {keybd_event.VK_Z, false},

		"A": {keybd_event.VK_A, true},
		"B": {keybd_event.VK_B, true},
		"C": {keybd_event.VK_C, true},
		"D": {keybd_event.VK_D, true},
		"E": {keybd_event.VK_E, true},
		"F": {keybd_event.VK_F, true},
		"G": {keybd_event.VK_G, true},
		"H": {keybd_event.VK_H, true},
		"I": {keybd_event.VK_I, true},
		"J": {keybd_event.VK_J, true},
		"K": {keybd_event.VK_K, true},
		"L": {keybd_event.VK_L, true},
		"M": {keybd_event.VK_M, true},
		"N": {keybd_event.VK_N, true},
		"O": {keybd_event.VK_O, true},
		"P": {keybd_event.VK_P, true},
		"Q": {keybd_event.VK_Q, true},
		"R": {keybd_event.VK_R, true},
		"S": {keybd_event.VK_S, true},
		"T": {keybd_event.VK_T, true},
		"U": {keybd_event.VK_U, true},
		"V": {keybd_event.VK_V, true},
		"W": {keybd_event.VK_W, true},
		"X": {keybd_event.VK_X, true},
		"Y": {keybd_event.VK_Y, true},
		"Z": {keybd_event.VK_Z, true},

		"0": {keybd_event.VK_0, false},
		"1": {keybd_event.VK_1, false},
		"2": {keybd_event.VK_2, false},
		"3": {keybd_event.VK_3, false},
		"4": {keybd_event.VK_4, false},
		"5": {keybd_event.VK_5, false},
		"6": {keybd_event.VK_6, false},
		"7": {keybd_event.VK_7, false},
		"8": {keybd_event.VK_8, false},
		"9": {keybd_event.VK_9, false},
		" ": {keybd_event.VK_SPACE, false},

		")": {keybd_event.VK_0, true},
		"!": {keybd_event.VK_1, true},
		"@": {keybd_event.VK_2, true},
		"#": {keybd_event.VK_3, true},
		"$": {keybd_event.VK_4, true},
		"%": {keybd_event.VK_5, true},
		"^": {keybd_event.VK_6, true},
		"&": {keybd_event.VK_7, true},
		"*": {keybd_event.VK_8, true},
		"(": {keybd_event.VK_9, true},

		"-":  {keybd_event.VK_MINUS, false},
		"=":  {keybd_event.VK_EQUAL, false},
		"[":  {keybd_event.VK_SP4, false},
		"]":  {keybd_event.VK_SP5, false},
		"\\": {keybd_event.VK_BACKSLASH, false},
		";":  {keybd_event.VK_SEMICOLON, false},
		"'":  {keybd_event.VK_SP7, false},
		",":  {keybd_event.VK_COMMA, false},
		".":  {keybd_event.VK_SP6, false},
		"/":  {keybd_event.VK_SLASH, false},
		"`":  {keybd_event.VK_GRAVE, false},

		"_":  {keybd_event.VK_MINUS, true},
		"+":  {keybd_event.VK_EQUAL, true},
		"{":  {keybd_event.VK_SP4, true},
		"}":  {keybd_event.VK_SP5, true},
		"|":  {keybd_event.VK_BACKSLASH, true},
		":":  {keybd_event.VK_SEMICOLON, true},
		"\"": {keybd_event.VK_SP7, true},
		"<":  {keybd_event.VK_COMMA, true},
		">":  {keybd_event.VK_SP10, true},
		"?":  {keybd_event.VK_SLASH, true},
		"~":  {keybd_event.VK_GRAVE, true},

		"ENTER":     {keybd_event.VK_ENTER, false},
		"TAB":       {keybd_event.VK_TAB, false},
		"BACKSPACE": {keybd_event.VK_DELETE, false},
	}
)

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
