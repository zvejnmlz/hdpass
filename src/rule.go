package hdpass

import (
	"crypto/rand"
	"fmt"
)

// Rule Chars map of ALPHABET:REQUIRED
type rule struct {
	len              uint8
	requiredChars    map[string]uint8
	requiredCharKeys []string
	optionalChars    string
	maxRequired      uint8
}

const RuleUpperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const RuleLowerAlpha = "abcdefghijklmnopqrstuvwxyz"
const RuleDigit = "0123456789"
const RuleMinus = "-"
const RuleUnderline = "_"
const RuleDot = "."
const RuleSpace = " "
const RuleSpecial1 = "!@#$%^&*=+~"
const RuleSpecial2 = "\\(){}[]<>;:,?/\"'"

func NewRule(len uint8, chars map[string]uint8) (*rule, error) {
	m := &rule{
		len:           len,
		requiredChars: make(map[string]uint8),
	}

	maxRequired := 0

	for k, v := range chars {
		if v != 0 {
			maxRequired += int(v)
			m.requiredChars[k] = v
			m.requiredCharKeys = append(m.requiredCharKeys, k)
		} else {
			m.optionalChars = m.optionalChars + k
		}
	}

	if maxRequired > int(len) {
		return nil, fmt.Errorf("required chars(%d) more len(%d)", maxRequired, len)
	}

	m.maxRequired = uint8(maxRequired)

	return m, nil
}

func (m *rule) Generate(vector string) string {
	if vector == "" {
		var buf = make([]byte, 32)
		_, _ = rand.Read(buf)
		vector = string(buf)
	}

	var output string
	var alphabetLen = len(m.optionalChars)
	var i uint8

	var r = NewReader(vector)

	for i = 0; i < m.len-m.maxRequired; i++ {
		p := r.IntRange(alphabetLen)
		output += m.optionalChars[p-1 : p]
	}

	for _, alphabet := range m.requiredCharKeys {
		count := m.requiredChars[alphabet]
		alphabetLen = len(alphabet)

		for i = 0; i < count; i++ {
			p := r.IntRange(alphabetLen)
			output += alphabet[p-1 : p]
		}
	}

	return r.Shuffle(output)
}

func (m *rule) Check(input string) error {
	//TODO
	//iLen := len(input)
	//if iLen > 255 || uint8(iLen) < m.len {
	//	return fmt.Errorf("minimum length must be >= %d characters", m.len)
	//}
	//
	//for alphabet, count := range m.requiredChars {
	//	var found uint8 = 0
	//
	//	for i, v := range alphabet {
	//		fmt.Sprintln(i, v)
	//		found++
	//	}
	//
	//	if found < count {
	//		return fmt.Errorf("required chars %s", alphabet)
	//	}
	//
	//}

	return nil
}
