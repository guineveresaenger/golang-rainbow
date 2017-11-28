package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	string := "Hello, Guinevere. You're a real dev now. No pressure...adding more chars to not run out of extra too much extra now not too much extra this can go on forrrrrevvvverrrrr"

	for i := 0; i < 50; i++ {
		rainbow(string, i)
	}

}

func rainbow(s string, lineCount int) {
	for i := 0; i < len(s); i++ {

		subindex := (i + lineCount) % 30 // will give the index of each 30 char substring and rotate through colors

		switch {
		case subindex >= 0 && subindex < 5:
			red(string(s[i]))
		case subindex >= 5 && subindex < 10:
			yellow(string(s[i]))
		case subindex >= 10 && subindex < 15:
			green(string(s[i]))
		case subindex >= 15 && subindex < 20:
			blue(string(s[i]))
		case subindex >= 20 && subindex < 25:
			magenta(string(s[i]))
		case subindex >= 25 && subindex < 30:
			cyan(string(s[i]))
		default:
			fmt.Printf(string(s[i]))
		}
	}
	fmt.Printf("\n")
}

func red(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.RedString(s), " "))
}

func yellow(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.YellowString(s), " "))
}

func blue(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.BlueString(s), " "))
}

func green(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.GreenString(s), " "))
}

func magenta(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.MagentaString(s), " "))
}

func cyan(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.CyanString(s), " "))
}
