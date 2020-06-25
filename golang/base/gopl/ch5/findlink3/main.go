package main

import "os"

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	return nil // TODO 真实逻辑
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
