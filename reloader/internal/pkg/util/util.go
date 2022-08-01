package util

type List []string

func (l *List) Contain(s string) bool {
	for _, v := range *l {
		if v == s {
			return true
		}
	}
	return false
}