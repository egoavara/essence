package must

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
func MustGet(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
