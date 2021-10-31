package util

//Distinct will return only unique value of an array
func Distinct(e []string) []string {
	r := []string{}

	for _, s := range e {
		if !contains(r[:], s) {
			r = append(r, s)
		}
	}
	return r
}

func contains(e []string, c string) bool {
	for _, s := range e {
		if s == c {
			return true
		}
	}
	return false
}
