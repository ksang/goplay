package main

import (
	"errors"
	"flag"
	"log"

	"github.com/ksang/goplay/quiz"
)

var (
	// quiz function name
	QUIZFUNC string
)

func init() {
	flag.StringVar(&QUIZFUNC, "quiz", "", "run code snippets in quiz e.g.:\"-quiz test\"")
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

func main() {
	flag.Parse()
	if len(QUIZFUNC) > 0 {
		if err := play("quiz", QUIZFUNC); err != nil {
			log.Fatal(err)
		}
		return
	}
	flag.PrintDefaults()
}
