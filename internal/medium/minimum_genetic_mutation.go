package medium

type Queue struct {
	storage []string
	current int
}

func NewQueue() Queue {
	return Queue{storage: []string{}, current: 0}
}

func (q *Queue) IsEmpty() bool {
	return q.current == len(q.storage)
}

func (q *Queue) Add(gen string) {
	q.storage = append(q.storage, gen)
}

func (q *Queue) Get() string {
	el := q.storage[q.current]
	q.current++
	return el
}
func (q *Queue) Has(gen string) bool {
	for _, v := range q.storage {
		if v == gen {
			return true
		}
	}
	return false
}

func MinMutation(start string, end string, bank []string) int {
	queue := NewQueue()
	queue.Add(start)
	counter := make([][]string, 0)

	for !queue.IsEmpty() {
		cur := queue.Get()
		for _, gene := range bank {
			if !queue.Has(gene) {
				if getDiff(cur, gene) == 1 {
					queue.Add(gene)
					counter = addCounter(counter, cur, gene)
					if gene == end {
						return getCounter(counter, start, end)
					}
				}
			}
		}
	}

	return -1
}

func addCounter(counter [][]string, cur string, gene string) [][]string {
	counter = append(counter, []string{cur, gene})
	for i, v := range counter {
		if v[len(v)-1] == cur {
			counter = append(counter, append(counter[i], gene))
		}
	}

	return counter
}

func getCounter(counter [][]string, start string, end string) int {
	for i := range counter {
		if counter[i][len(counter[i])-1] == end && counter[i][0] == start {
			return len(counter[i]) - 1
		}
	}
	return -1
}

func getDiff(start string, gene string) int {
	diff := 0
	for i := 0; i < 8; i++ {
		if start[i] != gene[i] {
			diff++
		}
	}

	return diff
}
