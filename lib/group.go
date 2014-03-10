package lib

type Group []string

func (g Group) Get(n int) string {
	strs := []string(g)
	if len(strs) >= n+1 {
		return strs[n]
	}
	return ""
}
