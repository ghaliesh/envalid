package utils

func Filter(strs []string, predicate func(s string) bool) []string {
	var result []string
	for _, s := range strs {
		match := predicate(s)
		if match {
			result = append(result, s)
		}
	}

	return result
}
