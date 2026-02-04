package main

import "sort"

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return [][]int{}
    }

    sort.Ints(nums)
    result := [][]int{}

    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        if nums[i] > 0 {
            break
        }

        left := i + 1
        right := len(nums) - 1
        target := -nums[i]

        for left < right {
            sum := nums[left] + nums[right]

            if sum == target {
                result = append(result, []int{nums[i], nums[left], nums[right]})

                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }

    return result
}