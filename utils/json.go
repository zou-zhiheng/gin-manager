package utils

// JsonDeleteOne 删除json字符串中的第一个匹配的字符串
func JsonDeleteOne(str string, target string) string {

	if str == "" {
		return ""
	}

	if target == "" {
		return str
	}

	for i := 0; i < len(str); i++ {
		flag := false
		for j := 0; j < len(target); j++ {
			if target[j] == str[i] {
				flag = true
				i++
			} else {
				break
			}
		}
		if flag {
			if str[i-1-len(target)] == ',' {
				return str[:i-1-len(target)] + str[i:]
			}
			return str[:i-len(target)] + str[i:]
		}
	}

	return str
}
