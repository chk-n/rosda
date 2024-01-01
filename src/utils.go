package dash

import "encoding/json"

func prependBytes(s []byte, v byte) []byte {
	s = append(s, v)
	copy(s[1:], s)
	s[0] = v
	return s
}

func structToStr(s any) (*string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	str := string(b)
	return &str, nil
}

func strToStruct(s string, d any) error {
	return json.Unmarshal([]byte(s), &d)
}
