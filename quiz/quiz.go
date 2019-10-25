package quiz

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Quiz interface {
	Run(funcname string) error
	Functions() map[string]interface{}
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
	start := time.Now()
	defer func() {
		fmt.Printf("=> function: %s \ttime elapsed: %s\n", funcname, time.Now().Sub(start))
	}()
	return f.(func() error)()
}

func (q *quiz) Functions() map[string]interface{} {
	return q.funcMap
}

func New() Quiz {
	q := quiz{
		funcMap: make(map[string]interface{}),
	}
	q.MustRegister("test", test)
	q.MustRegister("findfuncarg", runFindFuncArgument)
	q.MustRegister("rotateimg", runRotateImg)
	q.MustRegister("editdistance", runEditDistance)
	q.MustRegister("flattenbintree", runFlattenBinaryTreeToLinkedList)
	q.MustRegister("levelorderbintree", runbinTreeLevelOrder)
	return &q
}

func test() error {
	fmt.Println("Quiz Test Success")
	return nil
}
