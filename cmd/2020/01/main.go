package main

import (
	"fmt"
	"github.com/alfiejsmith/advent-of-code/internal/pkg/filereader"
)


func main() {

	nums := filereader.ReadIntFile("input")

	var target int64 = 2020

	PartOne:
		for i := 0; i < len(nums); i++ {
			for j := i; j < len(nums); j++ {
				if nums[i] + nums[j] == target {
					fmt.Println("Part One:", nums[i] * nums[j])
					break PartOne
				}
			}
		}

	PartTwo:
		for i := 0; i < len(nums); i++ {
			for j := i; j < len(nums); j++ {
				for k := j; k < len(nums); k++ {
					if nums[i] + nums[j] + nums[k] == target {
						fmt.Println("Part Two:", nums[i] * nums[j] * nums[k])
						break PartTwo
					}
				}
			}
		}

}
