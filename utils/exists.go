package utils

func Exists(m map[string]string, target string, lookForKeys bool) bool {

	for k, v := range m {
		if lookForKeys && k == target {
			return true
		}

		if !lookForKeys && v == target {
			return true
		}
	}

	return false
}
