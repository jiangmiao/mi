package mi

type String string

func (v String) Underscorize() String {
	var n = len(v)
	if n == 0 {
		return ""
	}
	us := make([]byte, n*2)
	var k, j int
	for j < n {
		if v[j] >= 'A' && v[j] <= 'Z' {
			if k > 0 {
				us[k] = '_'
				k++
			}
			us[k] = v[j] - 'A' + 'a'
		} else {
			us[k] = v[j]
		}
		k++
		j++
	}
	return String(us[:k])
}

func (s String) Camelize(args ...interface{}) String {
	var n = len(s)
	if n == 0 {
		return String("")
	}
	var capitalize bool
	if len(args) == 0 {
		capitalize = false
	} else {
		capitalize = args[0].(bool)
	}

	var r = make([]byte, n)
	var i, j int
	if capitalize {
		if 'a' <= s[i] && s[i] <= 'z' {
			r[j] = s[i] - 'a' + 'A'
			i++
			j++
		}
	}
	for j < n-1 {
		if s[j] == '_' {
			j++
			if 'a' <= s[j] && s[j] <= 'z' {
				r[i] = s[j] - 'a' + 'A'
			} else {
				r[i] = s[j]
			}
		} else {
			r[i] = s[j]
		}
		i++
		j++
	}
	if j < n {
		r[i] = s[j]
		i++
		j++
	}
	return String(r[:i])
}
