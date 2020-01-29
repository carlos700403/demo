package main

import "github.com/kidlj/demo/gopl.io/ch5/links"

import "log"

import "fmt"

import "os"

func main() {
	breadthFirst(os.Args[1:], crawl)
}

func breadthFirst(worklist []string, f func(item string) []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if _, ok := seen[item]; !ok {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	items, _, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return items
}