package main

import (
    "fmt"
    "math"
)

func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0

    for left < right {
        currentArea := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
        maxArea = int(math.Max(float64(maxArea), float64(currentArea)))

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}