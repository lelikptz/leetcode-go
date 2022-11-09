package medium

type Item struct {
	Price int
	Span  int
}

type StockSpanner struct {
	stack []Item
}

func Constructor() StockSpanner {
	return StockSpanner{stack: []Item{}}
}

func (s *StockSpanner) Next(price int) int {
	span := 1

	for i := len(s.stack) - 1; i >= 0; i-- {
		if s.stack[i].Price <= price {
			span += s.stack[i].Span
			i -= s.stack[i].Span - 1
		} else {
			break
		}
	}
	s.stack = append(s.stack, Item{price, span})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
