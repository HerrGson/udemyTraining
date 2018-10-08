package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Unsorted people:", studyGroup)

	sort.Slice(studyGroup, func(i, j int) bool { return studyGroup[i] < studyGroup[j] })
	fmt.Println("people sorted with sort.Slice :", studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Unsorted []string:", s)
	sort.Strings(s)
	fmt.Println("[]string sorted with sort.String", s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("Unsorted []int", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("[]int sorted with sort.Sort(sort.Reverse(sort.IntSlice(n)))", n)
}
