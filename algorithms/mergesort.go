package algorithms

/*
// Testing
import (
	"fmt"
)

func main() {
	nums := []int{1, 7, 9, 5, 3, 6, 4, 2, 49, 11, 8}
	Mergesort(nums)
	fmt.Println(nums)
}
*/
func Mergesort(nums []int) {
	temp := make([]int, len(nums))
	merge_sort(nums, temp, 0, len(nums)-1)
}

func merge_sort(nums, temp []int, leftStart, rightEnd int) {
	// case: 1 element
	if leftStart >= rightEnd {
		return
	}
	mid := (leftStart + rightEnd) / 2
	merge_sort(nums, temp, leftStart, mid)
	merge_sort(nums, temp, mid+1, rightEnd)
	merge_halves(nums, temp, leftStart, rightEnd)
}

func merge_halves(nums, temp []int, leftStart, rightEnd int) {
	// leftEnd is mid
	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1
	left, right, i := leftStart, rightStart, leftStart

	for left <= leftEnd && right <= rightEnd {
		if nums[left] <= nums[right] {
			temp[i] = nums[left]
			left++
		} else {
			temp[i] = nums[right]
			right++
		}
		i++
	}

	copy(temp[i:i+leftEnd-left+1], nums[left:leftEnd+1])
	copy(temp[i:i+rightEnd-right+1], nums[right:rightEnd+1])
	copy(nums[leftStart:rightEnd+1], temp[leftStart:rightEnd+1])
}
