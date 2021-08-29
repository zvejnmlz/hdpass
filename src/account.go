package hdpass

import (
	"strconv"
)

type account struct {
	seed       string
	domain     string
	loginRule  *rule
	passRule   *rule
	loginIndex int
	passIndex  int
}

func (m *account) generateLogin(index int) string {
	return m.loginRule.Generate(m.seed + strconv.Itoa(index))
}

func (m *account) SelectLogin(index int) string {
	return m.generateLogin(index)
}

func (m *account) CurrentLogin() (string, int) {
	return m.generateLogin(m.loginIndex), m.loginIndex
}

func (m *account) PrevLogin() string {
	m.loginIndex--

	return m.generateLogin(m.loginIndex)
}

func (m *account) NextLogin() string {
	m.loginIndex++

	return m.generateLogin(m.loginIndex)
}

func (m *account) generatePass(index int) string {
	currentLogin, _ := m.CurrentLogin()

	return m.passRule.Generate(m.seed + m.domain + strconv.Itoa(index) + currentLogin)
}

func (m *account) SelectPassword(index int) string {
	return m.generatePass(index)
}

func (m *account) CurrentPassword() (string, int) {
	return m.generatePass(m.passIndex), m.passIndex
}

func (m *account) NextPassword() string {
	m.passIndex++

	return m.generatePass(m.passIndex)
}

func (m *account) PrevPassword() string {
	m.passIndex--

	return m.generatePass(m.passIndex)
}
