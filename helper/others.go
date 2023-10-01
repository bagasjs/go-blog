package helper

func ChooseString(a string, b string) string {
	if len(a) != 0 {
		return a
	}

	return b
}
