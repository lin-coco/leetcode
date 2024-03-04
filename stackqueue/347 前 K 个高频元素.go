package stackqueue

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	ts := make(Tops, 0, len(m))
	for k, v := range m {
		heap.Push(&ts, ZSet{Member: k, Score: v})
	}
	heap.Init(&ts)
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, (heap.Pop(&ts).(ZSet)).Member)
	}
	return res
}

// 大顶堆
type ZSet struct {
	Member int
	Score  int
}

type Tops []ZSet

func (t *Tops) Len() int {
	return len(*t)
}
func (t *Tops) Less(a, b int) bool {
	return (*t)[a].Score < (*t)[b].Score
}
func (t *Tops) Swap(a, b int) {
	(*t)[a], (*t)[b] = (*t)[b], (*t)[a]
}
func (t *Tops) Push(x interface{}) {
	*t = append(*t, x.(ZSet))
}
func (t *Tops) Pop() interface{} {
	res := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return res
}
