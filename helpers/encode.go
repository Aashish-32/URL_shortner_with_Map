package helpers

func Base62Encode(num int) string {

	base62Char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if num == 0 {
		return string(base62Char[0])
	} else {
		if num < 0 {
			num = -num
		}

		base := len(base62Char)
		bytes := make([]byte, 0)

		for num > 0 {
			bytes = append(bytes, base62Char[num%base])
			num /= base
		}

		return string(bytes)
	}
}
