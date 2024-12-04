package util

type StrBuild struct {
	data  []rune
	empty bool
}

func NewStrBuild() *StrBuild {
	return &StrBuild{
		data:  []rune{},
		empty: true,
	}
}

func (it *StrBuild) Append(r rune) {
	it.data = append(it.data, r)
	it.empty = false
}

func (it *StrBuild) String() string {
	return string(it.data)
}

func (it *StrBuild) Flush() string {
    s := it.String()
    it.Reset()

    return s
}

func (it *StrBuild) Reset() {
	it.data = []rune{}
    it.empty = true
}

func (it *StrBuild) IsEmpty() bool {
    return it.empty
}
