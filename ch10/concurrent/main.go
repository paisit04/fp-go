package main

import (
	"concurrent/pkg"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	output := pkg.ConcurrentFilter(ints, func(i int) bool {
		return i%2 == 0
	}, 3)
	fmt.Printf("%v\n", output)

	output = pkg.ConcurrentMap(output, func(i int) int {
		return i * 2
	}, 2)
	fmt.Printf("%v\n", output)

	result := pkg.ConcurrentFMap(output, func(i int) string {
		return "number: " + strconv.Itoa(i)
	}, 2)
	fmt.Printf("%v\n", result)

	fmt.Printf("\n\n\n===pipelines===\n")
	generated := pkg.Generator(1, 2, 3, 4)
	filtered := pkg.FilterNode(generated, func(i int) bool {
		return i%2 == 0
	})
	mapped := pkg.MapNode(filtered, func(i int) int {
		return i * 2
	})
	collected := pkg.Collector(mapped)
	fmt.Printf("%v\n", collected)

	fmt.Println("\n\nchaining")
	out := pkg.ChainPipes(pkg.Generator(1, 2, 3, 4),
		pkg.CurriedFilterNode(func(i int) bool { return i%2 == 0 }),
		pkg.CurriedMapNode(func(i int) int { return i * i }))

	fmt.Println(out)

	fmt.Println("\n\nchain pipes 2")
	out2 := pkg.ChainPipes2[string](pkg.CurriedCat("./main.go"),
		pkg.CurriedFilterNode(func(s string) bool { return strings.Contains(s, "func") }),
		pkg.CurriedMapNode(func(i string) string { return "\nline contains func: " + i }))

	fmt.Printf("%v\n", out2)
}
