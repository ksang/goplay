package algo

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Algo interface {
	Run(funcname string) error
	Functions() map[string]interface{}
}

type algo struct {
	funcMap map[string]interface{}
}

func (a *algo) MustRegister(funcname string, f interface{}) {
	if _, exist := a.funcMap[funcname]; exist {
		log.Fatalf("register failed, function: %s already exist", funcname)
	}
	a.funcMap[funcname] = f
}

func (a *algo) Run(funcname string) error {
	f, ok := a.funcMap[funcname]
	if !ok {
		return errors.New(fmt.Sprintf("function: %s not found", funcname))
	}
	start := time.Now()
	defer func() {
		fmt.Printf("=> function: %s \ttime elapsed: %s\n", funcname, time.Now().Sub(start))
	}()
	return f.(func() error)()
}

func (a *algo) Functions() map[string]interface{} {
	return a.funcMap
}

func New() Algo {
	a := algo{
		funcMap: make(map[string]interface{}),
	}
	a.MustRegister("mergesort", runMergeSort)
	a.MustRegister("quicksort", runQuickSort)
	return &a
}
