package medium

type Char struct {
	x     int
	y     int
	value uint8
}

type StackItem struct {
	element Char
	index   int
	path    []Char
}

func (c Char) IsNearby(cur Char) bool {
	return (c.y == cur.y && (c.x == cur.x-1 || c.x == cur.x+1)) || (c.x == cur.x && (c.y == cur.y+1 || c.y == cur.y-1))
}

func Exist(board [][]byte, word string) bool {
	var cur *Char
	var stack []StackItem
	var path []Char

	for i := 0; i < len(word); i++ {
		chars := findChars(board, cur, word[i], path)
		if len(chars) == 0 && len(stack) == 0 {
			return false
		}

		for _, val := range chars {
			stack = append(stack, StackItem{
				element: val,
				index:   i,
				path:    path,
			})
		}

		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = &item.element
		i = item.index
		path = item.path
		path = append(path, *cur)
	}

	return true
}

func findChars(board [][]byte, cur *Char, value uint8, path []Char) []Char {
	var chars []Char
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if cur == nil {
				if value == board[y][x] {
					chars = append(chars, Char{
						x:     x,
						y:     y,
						value: board[y][x],
					})
				}
			} else if value == board[y][x] {
				c := Char{
					x:     x,
					y:     y,
					value: board[y][x],
				}
				if c.IsNearby(*cur) && !inPath(c, path) {
					chars = append(chars, c)
				}
			}
		}
	}

	return chars
}

func inPath(cur Char, path []Char) bool {
	for i := 0; i < len(path); i++ {
		if path[i].value == cur.value && path[i].x == cur.x && path[i].y == cur.y {
			return true
		}
	}

	return false
}
