package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(isPalindrome(131))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println(merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18}}))

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	//key表示 nums中的数字，value表示 出现次数
	countMap := map[int]int{}
	for i := range nums {
		value := nums[i]
		countMap[value] = countMap[value] + 1
	}
	for key, value := range countMap {
		if value == 1 {
			return key
		}
	}
	return -1
}

// 回文数字
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	//使用栈  遇到左括号入栈，遇到右括号判断栈顶元素是否匹配
	stack := []string{}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, string(char))
		} else {
			if len(stack) == 0 && (char == ')' || char == '}' || char == ']') {
				return false
			}

			if len(stack) == 0 {
				continue
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == ')' && top != "(") || (char == '}' && top != "{") || (char == ']' && top != "[") {
				return false
			}
		}
	}
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	commonPrefix := ""
	if len(strs) == 1 {
		commonPrefix = strs[0]
		return commonPrefix
	}
	k := 1
	for {
		flag := true
		for i := 1; i < len(strs); i++ {
			if k > len(strs[i]) || k > len(strs[i-1]) {
				flag = false
				break
			}
			prefix := strs[i][0:k]
			lastPrefix := strs[i-1][0:k]
			if prefix != lastPrefix {
				flag = false
				break
			}
		}
		if flag {
			commonPrefix = strs[0][0:k]
		} else {
			break
		}
		k++
	}
	return commonPrefix
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	isAddOne := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += isAddOne
		if digits[i] > 9 {
			digits[i] = 0
			isAddOne = 1
		} else {
			isAddOne = 0
			break
		}
	}
	if isAddOne == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				//判断是否有重复数字，有则删除
				//由于nums不会扩容，会修改原数组的值
				nums = append(nums[:j], nums[j+1:]...)
				j--
			}
		}
	}

	return len(nums)
}

// 合并区间
func merge(intervals [][]int) [][]int {
	//先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		item := intervals[i]
		nextItem := intervals[i+1]
		//重叠合并
		if nextItem[0] <= item[1] {
			mergeItem := make([]int, 2, 2)
			mergeItem[0] = item[0]
			end := nextItem[1]
			if item[1] > nextItem[1] {
				end = item[1]
			}
			mergeItem[1] = end
			//替换并删除下标n+1
			intervals[i] = mergeItem
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}
	return intervals
}

// 给定一个整数数组 nums 和一个整数目标值 target
// 请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
