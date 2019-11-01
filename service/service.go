package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Service interface {
	Run(funcname string) error
	Names() map[string]interface{}
}

type service struct {
	serviceMap map[string]interface{}
	path       string
	bindaddr   string
}

func (s *service) MustRegister(serviceName string, f interface{}) {
	if _, exist := s.serviceMap[serviceName]; exist {
		log.Fatalf("register failed, service: %s already exist", serviceName)
	}
	s.serviceMap[serviceName] = f
}

func (s *service) Run(serviceName string) error {
	f, ok := s.serviceMap[serviceName]
	if !ok {
		return errors.New(fmt.Sprintf("service: %s not found", serviceName))
	}
	start := time.Now()
	defer func() {
		fmt.Printf("=> service: %s \ttime elapsed: %s\n", serviceName, time.Now().Sub(start))
	}()
	http.Handle(s.path, f.(http.Handler))
	fmt.Printf("=> service: %s \tlistening at: %s%s\n", serviceName, s.bindaddr, s.path)
	return http.ListenAndServe(s.bindaddr, nil)
}

func (s *service) Names() map[string]interface{} {
	return s.serviceMap
}

func New() Service {
	s := service{
		serviceMap: make(map[string]interface{}),
		path:       "/service",
		bindaddr:   ":8080",
	}
	s.MustRegister("echo", &echo{})
	s.MustRegister("md5", &md5Service{4})
	return &s
}
