package hard

type MedianFinder struct {
	nums []int
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{nums: []int{}}
}

func (m *MedianFinder) AddNum(num int) {
	if len(m.nums) == 0 || m.nums[len(m.nums)-1] <= num {
		m.nums = append(m.nums, num)
	} else {
		for index, value := range m.nums {
			if value > num {
				m.nums = append(m.nums[:index+1], m.nums[index:]...)
				m.nums[index] = num
				break
			}
		}
	}
}

func (m *MedianFinder) FindMedian() float64 {
	key := len(m.nums) / 2
	if len(m.nums)%2 == 1 {
		return float64(m.nums[key])
	}

	return float64(m.nums[key]+m.nums[key-1]) / 2
}
