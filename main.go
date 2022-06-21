package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Println("EX: go run . something standard --align=right")
		return
	}
	if os.Args[2] != "shadow" && os.Args[2] != "standard" && os.Args[2] != "thinkertoy" {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Println("EX: go run . something standard --align=right")
		return
	}
	if strings.Contains(os.Args[3], "--align=") {
	} else {
		fmt.Println("Usage: go run . [STRING] [OPTION] [LETTERS]")
		fmt.Println("EX: go run . something --color=<color> so")
		return
	}
	if len(os.Args) == 4 {
		var input string = os.Args[1]
		draw(input)
	}
}

func draw(input string) {
	w := getWidth()
	var line1 string
	var line2 string
	var line3 string
	var line4 string
	var line5 string
	var line6 string
	var line7 string
	var line8 string
	var s []string
	var runeAlphabet []string
	var hasNewline bool
	var simpleNewLine bool
	var newInput string
	var newVal string
	var count int
	hasNewline = hasNewline
	file, err := os.Open(os.Args[2] + ".txt")
	err = err
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	for i := 0; i < len(input); i++ {
		if string(input[i]) == "\\" && i < len(input)-1 && string(input[i+1]) == "n" {
			if i+2 != len(input) {
				newInput += string(input[i])
				for j := i + 2; j < len(input); j++ {
					newVal += string(input[j])
					count++
				}
				for k := 0; k < count; k++ {
					input = input[:len(input)-1]
				}
			} else {
				simpleNewLine = true
			}
		}
	}
	a := strings.Contains(input, "\\n")
	if a {
		input = strings.ReplaceAll(input, "\\n", "")
		if simpleNewLine != true {
			hasNewline = true
		}
	}

	for i := 0; i <= 95; i++ {
		runeAlphabet = append(runeAlphabet, string(i+32))
	}
	var space bool
	var spacePos int
	var words int
	space = space
	spacePos = spacePos
	words = words
	for i := 0; i < len(input); i++ {
		if string(input[i]) == " " {
			words++
		}
	}
	for j := 0; j < len(input); j++ {
		for k := 0; k < len(runeAlphabet); k++ {
			if string(input[j]) == runeAlphabet[k] {
				for l := 1; l < 9; l++ {
					if l == 1 {
						line1 += s[k*9+l]
					}
					if l == 2 {
						line2 += s[k*9+l]
					}
					if l == 3 {
						line3 += s[k*9+l]
					}
					if l == 4 {
						line4 += s[k*9+l]
					}
					if l == 5 {
						line5 += s[k*9+l]
					}
					if l == 6 {
						line6 += s[k*9+l]
					}
					if l == 7 {
						line7 += s[k*9+l]
					}
					if l == 8 {
						line8 += s[k*9+l]
					}
				}
				if string(input[j]) == " " && os.Args[3] == "--align=justify" {
					w2 := int(w / 2)
					w2 = int(w)
					w2 = w2 / len(input) * 10
					if words <= 2 {
						w2 = w2 / (words * (words + 1))
					} else {
						for i := 1; i < words; i++ {
							w2 = w2 / (words * (words * i))
						}
					}

					for i := 0; i < w2; i++ {
						line1 += " "
					}
					for i := 0; i < w2; i++ {
						line2 += " "
					}
					for i := 0; i < w2; i++ {
						line3 += " "
					}
					for i := 0; i < w2; i++ {
						line4 += " "
					}
					for i := 0; i < w2; i++ {
						line5 += " "
					}
					for i := 0; i < w2; i++ {
						line6 += " "
					}
					for i := 0; i < w2; i++ {
						line7 += " "
					}
					for i := 0; i < w2; i++ {
						line8 += " "
					}
					space = true
					spacePos = j
				}
			}
		}
	}
	if len(line8) == 0 {
		for i := 0; i < len(line1); i++ {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	if simpleNewLine {
		fmt.Print("\n")
	}
	w2 := int(w / 2)
	if os.Args[3] == "--align=left" {
		w2 = 0
	} else if os.Args[3] == "--align=right" {
		w2 = int(w)
	} else if os.Args[3] == "--align=center" && words == 0 {
		w2 = int(w / 15 * 11)
	} else if os.Args[3] == "--align=center" && words >= 1 {
		w2 = int(w / 15 * 14)
	} else if os.Args[3] == "--align=justify" {
		w2 = int(w)
		w2 = w2 / len(input) * 10
	}
	if os.Args[3] != "--align=justify" {
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line1+"\n"))
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line2+"\n"))
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line3+"\n"))
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line4+"\n"))
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line5+"\n"))
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line6+"\n"))
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line7+"\n"))
		fmt.Printf(fmt.Sprintf("%%-%ds", w2), fmt.Sprintf(fmt.Sprintf("%%%ds", w2), line8+"\n"))
	} else {
		fmt.Println(line1)
		fmt.Println(line2)
		fmt.Println(line3)
		fmt.Println(line4)
		fmt.Println(line5)
		fmt.Println(line6)
		fmt.Println(line7)
		fmt.Println(line8)
	}
	if hasNewline {
		draw(newVal)
	}
}

func remove(a []string, i int) []string {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

// 1 à 855
