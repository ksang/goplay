package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/ksang/goplay/quiz"
)

var (
	// quiz function name
	QUIZFUNC string
	// list all available functions
	LIST bool
)

func init() {
	flag.StringVar(&QUIZFUNC, "quiz", "", "run code snippets in quiz e.g.:\"-quiz test\"")
	flag.BoolVar(&LIST, "list", false, "list all available functions")
}

func play(dir string, funcName string) error {
	switch dir {
	case "quiz":
		q := quiz.New()
		return q.Run(funcName)
	default:
		return errors.New("directory not found")
	}
	return nil
}

func list() {
	fmt.Println("quiz:")
	q := quiz.New()
	for k, _ := range q.Functions() {
		fmt.Printf("\t%s\n", k)
	}
}

func main() {
	flag.Parse()
	if LIST {
		list()
		return
	}
	if len(QUIZFUNC) > 0 {
		if err := play("quiz", QUIZFUNC); err != nil {
			log.Fatal(err)
		}
		return
	}
	flag.PrintDefaults()
}
