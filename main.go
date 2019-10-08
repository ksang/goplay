package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/ksang/goplay/algo"
	"github.com/ksang/goplay/quiz"
)

var (
	// quiz function name
	QUIZFUNC string
	// algo function name
	ALGOFUNC string
	// list all available functions
	LIST bool
)

func init() {
	flag.StringVar(&QUIZFUNC, "quiz", "", "run code snippets in quiz e.g.:\"-quiz test\"")
	flag.StringVar(&ALGOFUNC, "algo", "", "run code snippets in algo e.g.:\"-algo mergesort\"")
	flag.BoolVar(&LIST, "list", false, "list all available functions")
}

func play(dir string, funcName string) error {
	switch dir {
	case "quiz":
		q := quiz.New()
		return q.Run(funcName)
	case "algo":
		a := algo.New()
		return a.Run(funcName)
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
	fmt.Println("algo:")
	a := algo.New()
	for k, _ := range a.Functions() {
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
	if len(ALGOFUNC) > 0 {
		if err := play("algo", ALGOFUNC); err != nil {
			log.Fatal(err)
		}
		return
	}
	flag.PrintDefaults()
}
