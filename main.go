package main

/*
 * Count unique lines.  Dangerously!
 *
 * This loads each unique line into memory.  If there are a bunch of the same lines, this may be more efficient
 * than 'cat | sort | uniq -c'.  But if there are a lot of unique lines, this could be terrible.
 */

import (
	"os"
		"bufio"
	"fmt"
)

////////////////////////////////////////////
// Where the real stuff happens
////////////////////////////////////////////

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string]int)

	for scanner.Scan() {
		s := scanner.Text()

		m[s] = m[s] + 1
	}

	for line, count := range m {
		fmt.Printf("%d: %s\n", count, line)
	}
}
