package brackets

const testVersion = 4

// BracketStack ... might be overkill :)
type BracketStack []rune

// Empty ..
func (bs BracketStack) Empty() bool { return len(bs) == 0 }

// Push ..
func (bs *BracketStack) Push(r rune) { *bs = append(*bs, r) }

// Pop ..
func (bs *BracketStack) Pop() rune {
	l := len(*bs)
	if l == 0 {
		return 0
	}

	r := (*bs)[l-1:]
	*bs = (*bs)[:l-1]
	return r[0]
}

func matchingBrackets(opening, closing rune) bool {
	return (opening == '(' && closing == ')') ||
		(opening == '[' && closing == ']') ||
		(opening == '{' && closing == '}')
}

// Bracket ..
func Bracket(input string) (bool, error) {
	stack := new(BracketStack)

	for _, r := range input {
		if r == '(' || r == '[' || r == '{' {
			stack.Push(r)
		} else if !matchingBrackets(stack.Pop(), r) {
			return false, nil
		}
	}

	if !stack.Empty() {
		return false, nil
	}
	return true, nil
}
