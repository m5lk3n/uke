package common

func ReplaceAtIndex(s string, i int, r string) string {
	return s[:i] + r + s[i+1:]
}
