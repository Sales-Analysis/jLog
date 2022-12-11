package dotenv

// Contains returns a true if the slice contains the value.
func contains(value string, values []string) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}
	return false
}
