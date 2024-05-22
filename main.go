package main

func main() {
	array := [5]int{1, 2, 3, 5, 6}

	Constructor(array)
}

type NumArray struct {
	prefixTable []int
}

func Constructor(nums [5]int) NumArray {
	numArr := new(NumArray)
	numArr.prefixTable = make([]int, len(nums))
	numArr.prefixTable[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		numArr.prefixTable[i] = nums[i] + numArr.prefixTable[i-1]
	}
	return *numArr
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixTable[right]
	} else {
		return this.prefixTable[right] - this.prefixTable[left-1]
	}
}
