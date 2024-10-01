package problem0380

import "math/rand"

type RandomizedSet struct {
	values map[int]bool
	nums   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: make(map[int]bool),
		nums:   []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.values[val] {
		return false
	}

	this.values[val] = true
	this.nums = append(this.nums, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.values[val] {
		delete(this.values, val)
		for i, num := range this.nums {
			if num == val {
				this.nums = append(this.nums[:i], this.nums[(i+1):]...)
			}
		}
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
