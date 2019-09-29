package quiz

import (
	"errors"
	"fmt"
	"log"
)

type Quiz interface {
	Run(funcname string) error
}

type quiz struct {
	funcMap map[string]interface{}
}

func (q *quiz) MustRegister(funcname string, f interface{}) {
	if _, exist := q.funcMap[funcname]; exist {
		log.Fatalf("register failed, function: %s already exist", funcname)
	}
	q.funcMap[funcname] = f
}

func (q *quiz) Run(funcname string) error {
	f, ok := q.funcMap[funcname]
	if !ok {
		return errors.New(fmt.Sprintf("function: %s not found", funcname))
	}
	return f.(func() error)()
}

func New() Quiz {
	q := quiz{
		funcMap: make(map[string]interface{}),
	}
	q.MustRegister("test", test)
	return &q
}

func test() error {
	fmt.Println("Quiz Test Success")
	return nil
}
