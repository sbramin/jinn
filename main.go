package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

type jinn struct {
	scanner *bufio.Scanner
	wishes  map[string]string
	options map[string]bool
}

func main() {
	j := newJinn()
	//

	for c, _ := range j.wishes {
		j.giveChoice(c)
	}
	fmt.Println()
	for c, _ := range j.options {
		j.giveOption(c)
	}

	//
	fmt.Println(j.wishes)
	fmt.Println(j.options)
}

func (j *jinn) giveChoice(choice string) {
	fmt.Printf("What do you wish your %v to be, it is currently %v: ", color.GreenString(choice), color.MagentaString(j.wishes[choice]))
	j.scanner.Scan()
	if j.scanner.Err() != nil {
		log.Println(j.scanner.Err())
	}
	if j.scanner.Text() != "" {
		j.wishes[choice] = j.scanner.Text()
	}

}

func (j *jinn) giveOption(option string) {
	fmt.Printf("Do you wish to add a %v, answer %v: ", color.GreenString(option), color.MagentaString("Yes/No"))
	j.scanner.Scan()
	if j.scanner.Err() != nil {
		log.Println(j.scanner.Err())
	}
	j.options[option] = yes(j.scanner.Text())

}

func newJinn() *jinn {
	var j jinn
	j.scanner = bufio.NewScanner(os.Stdin)
	j.wishes = make(map[string]string)

	j.wishes["application name"] = "test-app"
	j.wishes["application description"] = "this is a test application"
	j.wishes["team"] = "billing"
	j.wishes["cloud"] = "aws"

	j.options = make(map[string]bool)
	j.options["kafka sink"] = false
	j.options["kafka source"] = false
	j.options["grpc api"] = false
	j.options["grpc client"] = false
	j.options["mongo db"] = false
	j.options["sql db"] = false

	return &j
}

func yes(i string) bool {
	i = strings.ToLower(i)
	return strings.Contains(i, "y")
}
