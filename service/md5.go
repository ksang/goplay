package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type md5Service struct {
	numGoroutine int
}

type result struct {
	path string
	hash string
}

func md5Worker(ctx context.Context, wg *sync.WaitGroup, wq chan string, hc chan result) {
	defer wg.Done()
	for path := range wq {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("worker error encountered, path: %s error: %v\n", path, err)
			return
		}
		hc <- result{
			path: path,
			hash: fmt.Sprintf("%x", md5.Sum(data)),
		}
		select {
		case <-ctx.Done():
			fmt.Println("operation canceled, worker quiting")
			return
		default:
			continue
		}
	}
}

func (m *md5Service) getMD5ResultsConcurrent(ctx context.Context, root string) (map[string]string, error) {
	var wg sync.WaitGroup
	res := make(map[string]string)
	goroutine := m.numGoroutine
	workQueue := make(chan string, goroutine)
	hashc := make(chan result, goroutine)
	errc := make(chan error)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg.Add(goroutine)
	for i := 0; i < goroutine; i++ {
		go md5Worker(ctx, &wg, workQueue, hashc)
	}

	go func() {
		for r := range hashc {
			res[r.path] = r.hash
		}
	}()

	go func() {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case workQueue <- path:
			case <-ctx.Done():
				fmt.Println("request canceled")
				return errors.New("request canceled")
			}
			return nil
		})
		close(workQueue)
		errc <- err
	}()
	defer close(hashc)
	wg.Wait()
	err := <-errc
	return res, err
}

func (m *md5Service) getMD5Results(ctx context.Context, root string) (map[string]string, error) {
	res := make(map[string]string)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		res[path] = fmt.Sprintf("%x", md5.Sum(data))
		return nil
	})
	return res, err
}

func (m *md5Service) writeResults(ctx context.Context, w http.ResponseWriter, res map[string]string) {
	r := ""
	for path, value := range res {
		r += fmt.Sprintf("%s  -  %s\n", value, path)
	}
	w.Write([]byte(r))
}

func (m *md5Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form from client: %s error: %v", r.UserAgent(), err)
		w.WriteHeader(502)
		fmt.Fprintf(w, "failed to parse form\n")
		return
	}
	path := r.Form.Get("path")
	if path == "" {
		return
	}
	res, err := m.getMD5ResultsConcurrent(r.Context(), path)
	if err != nil {
		log.Printf("failed to get md5 from client: %s path: %s error: %v", r.UserAgent(), path, err)
		w.WriteHeader(502)
		fmt.Fprintf(w, "failed to get md5 from path: %s\n", path)
		return
	}
	m.writeResults(r.Context(), w, res)
}
