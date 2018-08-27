package paren

const (
	PAREN_ROUND = iota
	PAREN_SQUARE
	PAREN_BIG
)

type Context struct {
	parentype int8
	prev      *Context
}

func isValid(s string) bool {
	this := new(Context)
	for _, char := range s {
		switch char {
		case '(':
			this = &Context{
				parentype: PAREN_ROUND,
				prev:      this,
			}
		case ')':
			if this.parentype != PAREN_ROUND || this.prev == nil {
				return false
			}
			this = this.prev
		case '[':
			this = &Context{
				parentype: PAREN_SQUARE,
				prev:      this,
			}
		case ']':
			if this.parentype != PAREN_SQUARE || this.prev == nil {
				return false
			}
			this = this.prev
		case '{':
			this = &Context{
				parentype: PAREN_BIG,
				prev:      this,
			}
		case '}':
			if this.parentype != PAREN_BIG || this.prev == nil {
				return false
			}
			this = this.prev
		}
	}
	return this.prev == nil
}
