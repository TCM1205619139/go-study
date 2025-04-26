package pkg

func Maybe(a *interface{}, b *interface{}) *interface{} {
	if a == nil {
		return b
	}
	return a
}

func MaybeString(a string, b string) string {
	if a == "" {
		return b
	}
	return a
}
