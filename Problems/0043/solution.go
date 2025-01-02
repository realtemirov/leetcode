package problem0043


func Multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    l1, l2 := len(num1), len(num2)

	bytes := make([]byte, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
            v := (num1[i]-'0')*(num2[j]-'0')
            bytes[i+j+1] += v
            if bytes[i+j+1] >= 10 {
                bytes[i+j] += bytes[i+j+1]/10
                bytes[i+j+1] %= 10
            }
		}
	}

    for bytes[0] == 0 {
        bytes = bytes[1:]
    }

    for i := range bytes {
        bytes[i] += '0'
    }

	return string(bytes)
}