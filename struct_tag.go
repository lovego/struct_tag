package struct_tag

func Get(tag, key string) string {
	value, _ := Lookup(tag, key)
	return value
}

func Lookup(tag, key string) (value string, ok bool) {
	for tag != "" {
		var name, quotedValue string
		name, quotedValue, tag = stripNameValuePair(tag)
		if name == "" || quotedValue == "" {
			break
		}
		if name == key {
			value, err := Unquote(quotedValue)
			if err != nil {
				break
			}
			return value, true
		}
	}
	return "", false
}

func Parse(tag string) (result map[string]string) {
	result = make(map[string]string)
	for tag != "" {
		var name, quotedValue string
		name, quotedValue, tag = stripNameValuePair(tag)
		if name == "" || quotedValue == "" {
			break
		}
		value, err := Unquote(quotedValue)
		if err != nil {
			break
		}
		result[name] = value
	}
	return result
}

func stripNameValuePair(tag string) (string, string, string) {
	if tag = trimLeadingSpace(tag); tag == "" {
		return "", "", ""
	}
	var name, quotedValue string
	if name, tag = stripName(tag); name == "" || tag == "" {
		return "", "", ""
	}
	quotedValue, tag = stripValue(tag)
	return name, quotedValue, tag
}

func trimLeadingSpace(tag string) string {
	i := 0
	for ; i < len(tag); i++ {
		switch tag[i] {
		case ' ', '\n', '\r', '\t':
		default:
			return tag[i:]
		}
	}
	return tag[i:]
}

func stripName(tag string) (string, string) {
	for i := 0; i < len(tag); i++ {
		// Scan to colon.
		if tag[i] == ':' {
			if i+1 < len(tag) {
				return tag[:i], tag[i+1:]
			} else {
				return tag[:i], ""
			}
		}
		// A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		if tag[i] <= ' ' || tag[i] == '"' || tag[i] == 0x7f {
			return "", ""
		}
	}
	return "", ""
}

func stripValue(tag string) (string, string) {
	// Scan quoted string.
	if tag[0] != '"' {
		return "", ""
	}
	for i := 1; i < len(tag); i++ {
		switch tag[i] {
		case '"':
			return tag[:i+1], tag[i+1:]
		case '\\':
			i++
		}
	}
	return "", ""
}
