package ppaster

func ConvertString(str string) string {

	bs := []byte(str)

	for i, r := range bs {
		if r == 10 {
			s := append(bs[:i], byte(32))
			bs = append(s, bs[i+1:]...)
		}
	}

	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == 32 && bs[i+1] == 32 {
			bs = append(bs[:i], bs[i+1:]...)
		}
	}

	s := string(bs)

	return s
}
