package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = FakeSearch("web", "The Go Programming Language", "http://golang.org")
	Image = FakeSearch("image", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Video = FakeSearch("video", "Concurrency is not Parallelism", "https://www.youtube.com/watch?v=cN_DpYBzKso")
)

type Result struct {
	Title, URL string
}

type SearchFunc func(query string) Result

func Search(query string) ([]Result, error) {
	results := []Result{
		Web(query),
		Image(query),
		Video(query),
	}

	return results, nil
}

func FakeSearch(kind, title, url string) SearchFunc {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Title: fmt.Sprintf("%s(%q): %s", kind, query, title),
			URL:   url,
		}
	}
}

func main() {
	start := time.Now()
	results, err := Search("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed, err)
}
