package utils

func IsValueInMap(list []string, search string) bool {
	for _, v := range list {
		if v == search {
			return true
		}
	}
	return false
}

func AreStringArrayEqual(one []string, two []string) bool {
	if len(one) != len(two) {
		return false
	}

	for _, v := range one {
		if IsValueInMap(two, v) == false {
			return false
		}
	}

	return true
}
