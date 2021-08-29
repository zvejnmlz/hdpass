package hdpass

import (
	"fmt"
)

type seed struct {
	seed    string
	domains map[string]*account
}

func NewSeed(seedRule *rule) (*seed, error) {
	s, err := RestoreFromSeed(seedRule.Generate(""), seedRule)

	return s, err
}

func RestoreFromSeed(input string, seedRule *rule) (*seed, error) {
	if err := seedRule.Check(input); err != nil {
		return nil, fmt.Errorf("seed validation (%s)", err)
	}

	return &seed{
		seed:    input,
		domains: make(map[string]*account),
	}, nil
}

func (m *seed) CurrentSeed() string {
	return m.seed
}

func (m *seed) NewDomain(domain string, loginRule *rule, passRule *rule) *account {
	l := &account{
		loginRule: loginRule,
		passRule:  passRule,
		seed:      m.seed,
		domain:    domain,
	}

	m.domains[domain] = l

	return l
}

func (m *seed) CurrentDomain() string {
	return m.CurrentDomain()
}

func (m *seed) SelectDomain(domain string) (*account, error) {
	var l *account
	var ok bool

	if l, ok = m.domains[domain]; !ok {
		return nil, fmt.Errorf("")
	}

	return l, nil
}
