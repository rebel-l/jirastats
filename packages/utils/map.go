package utils

func IsValueInMap(list []string, search string) bool {
	for _, v := range list {
		if v == search {
			return true
		}
	}
	return false
}
