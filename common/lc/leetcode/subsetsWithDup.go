package leetcode

import (
	"fmt"
	"sort"
	"strconv"
)

func SubsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	if len(nums) == 1 {
		return [][]int{nums, []int{}}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	d := make(map[string]struct{})
	n := [][]int{}
	n = Fk(nums[0], nums[1:], &d)
	n = append(n, []int{})
	return n
}

func Fk(firstNum int, nums []int, diffMap *map[string]struct{}) (allNum [][]int) {
	fmt.Println("fn", firstNum, "arr", nums, "diff", diffMap, "arrM")
	if len(nums) == 0 {
		a := []int{firstNum}
		(*diffMap)[getKey(a)] = struct{}{}
		allNum = append(allNum, a)
		//fmt.Println("len0", allNum)
		return
	}
	allNum = Fk(nums[0], nums[1:], diffMap)
	lN := len(allNum)
	for i := 0; i < lN; i++ {
		lastK := getKey(allNum[i])
		newKey := lastK + "-" + strconv.Itoa(firstNum)
		_, ok := (*diffMap)[newKey]
		if ok {
			continue
		}
		x := make([]int, len(allNum[i]))
		copy(x, allNum[i])
		newArr := append(x, firstNum)
		(*diffMap)[newKey] = struct{}{}
		allNum = append(allNum, newArr)
		//fmt.Println("before insert", firstNum, allNum[i])
		//fmt.Println("for all num", allNum)
	}
	a := []int{firstNum}
	k := getKey(a)
	if _, ok := (*diffMap)[k]; !ok {
		(*diffMap)[k] = struct{}{}
		allNum = append(allNum, a)
		//fmt.Println("last all num", allNum)
	}
	return
}

func getKey(arr []int) (k string) {
	for i := 0; i < len(arr); i++ {
		k = k + strconv.Itoa(arr[i]) + "-"
	}
	return
}

/**
  1 2
 1  2  12
1 2 3

1 2 12 13 23 123 3
 1234

1 2 12 13 23 123 3 14 24 124 134 234 1234 4

2

2 2
2 22

2 2 2

2 22 222

*/
