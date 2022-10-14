package gstr

const (
	NotFoundIndex = -1
)

// SearchArray 查询子串在字符串列表里的索引,不存在则返回-1
func SearchArray(strList []string, s string) int {
	for i, v := range strList {
		if s == v {
			return i
		}
	}
	return NotFoundIndex
}

// InArray 查询子串是否在字符串列表里
func InArray(a []string, s string) bool {
	return SearchArray(a, s) != NotFoundIndex
}
