package berrors

func Unwrap[T interface{}](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
