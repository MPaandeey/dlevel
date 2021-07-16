package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func regx(input string) string {

	ilevel, _ := strconv.Atoi(input)

	prefix := "("
	center := "\\w*\\."
	suffix := "\\w*$)"
	fmiddle := strings.Repeat(center, ilevel)
	pattern := prefix + fmiddle + suffix

	return pattern

}

func scanner() {
	lines := make(map[string]bool)

	re := regexp.MustCompile(regx(os.Args[1]))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		result := re.FindAllString(scanner.Text(), -1)
		for _, domain := range result {

			//Remove duplicates from result
			if lines[domain] {
				continue
			}
			lines[domain] = true
			fmt.Println(domain)
		}
	}
}

func main() {
	info, _ := os.Stdin.Stat()
	if len(os.Args) > 1 && info.Mode()&os.ModeNamedPipe != 0 {
		
		input, _ := strconv.Atoi(os.Args[1])
		if input > 0 {
			scanner()
		}
	}
}
