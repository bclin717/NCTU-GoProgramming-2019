package main

import "fmt"

// Implement a simple int set with array/slice

type IntSet struct {
	nums []int
}

func (set *IntSet) Add(num int) {
	for _, n := range set.nums {
		if num == n {
			return
		}
	}
	set.nums = append(set.nums, num)
}

func (set IntSet) Print() {
	fmt.Print("{ ")
	for _, n := range set.nums {
		fmt.Printf("%d ", n)
	}
	fmt.Print("}\n")
}

func main() {
	set := IntSet{}
	set.Add(12)
	set.Add(555)
	set.Add(12)
	set.Add(555)
	set.Print()
}
