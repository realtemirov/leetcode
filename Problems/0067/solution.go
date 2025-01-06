package problem0067

func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	bin := make([]byte, len(a)+1)
	var (
		carry byte
        a_num byte
		b_num byte
	)
	for i, j := len(a), len(b); i >= 1 || j >= 1; {
		i, j = i-1, j-1

		if i >= 0 {
			a_num = a[i] - '0'
		} else {
            a_num = 0
        }
		if j >= 0 {
			b_num = b[j] - '0'
		} else {
            b_num = 0
        }
		num := a_num ^ b_num ^ carry
		carry = a_num & b_num | carry & (a_num ^ b_num)
		bin[i+1] = num + '0'
	}
	if carry == 1 {
		bin[0] = '1'
		return string(bin)
	}
	return string(bin[1:])
}