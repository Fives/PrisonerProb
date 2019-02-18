package main

import (
	"fmt"
)

func main() {
	var freedp, allfreep, freedf, allfreef, next int
	var arr, arrorder, arrfield []int
	var perm [][]int
	//var winp, winf [][]int
	for k := 2; k < 7; k++ {
		arr = makeRange(k)
		perm = permutations(arr)
		allfreep, allfreef = 0, 0
		//winp, winf = nil, nil
		for _, r := range perm {
			freedp, freedf = 0, 0
			arrorder = nil
			for i := range r { // i = prisoner number; v = number in box of prisoner's number
				if freedp != i {
					continue
				}
				if i < len(r)/2 {
					arrorder = r[i:]
				} else {
					arrorder = append(r[i:], r[:i]...) // start slice at prisoner number
				}
				for _, num := range arrorder[:len(r)/2] {
					if num == i { // prisoner found his number
						freedp++
						//fmt.Printf("It took prisoner #%d %d tries to find his number\n", i, tries)
					}
				}
				if freedp == len(r) { // if all prisoners fonud his number, then all free
					allfreep++
					//winp = append(winp, r)
				}
			}
			for i, v := range r { // i = prisoner number; v = number in box of prisoner's number
				if freedf != i {
					continue
				}
				next = v
				arrfield = nil
				//fmt.Printf("Initial ordered field: %v\n", arrfield)
				for r[next] != i {
					next = r[next]
					arrfield = append(arrfield, next)
				}
				if len(arrfield)+1 < len(r)/2 {
					freedf++
				}
				if freedf == len(r) { // if all prisoners fonud his number, then all free
					allfreef++
					//winf = append(winf, r)
				}
			}
		}
		percentpermute := float64(allfreep) / float64(len(perm))
		percentfield := float64(allfreef) / float64(len(perm))
		fmt.Printf("#prisoners: %d\t#permutations: %d\n", k*2, len(perm))
		fmt.Printf("Times found -  permute: %v\tfield: %v\n", allfreep, allfreef)
		fmt.Printf("Percent found -  permute: %v\tfield: %v\n", percentpermute, percentfield)
		//fmt.Printf("Winning permute: %v\tfield: %v\n", winp, winf)
	}
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func makeRange(num int) []int {
	a := make([]int, num*2)
	for i := range a {
		a[i] = i
	}
	return a
}
