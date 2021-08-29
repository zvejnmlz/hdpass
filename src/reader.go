package hdpass

import (
	crand "crypto/rand"
	"crypto/sha256"
	"hash/crc64"
	"math/big"
	"math/rand"
)

type reader struct {
	init       string
	position   int
	vector     []byte
	nextVector []byte
}

func NewReader(vector string) *reader {
	h := sha256.Sum256([]byte(vector + "InitSalt"))
	nh := sha256.Sum256(h[:])

	v := &reader{
		init:       vector,
		vector:     h[:],
		nextVector: nh[:],
	}

	return v
}

func (m *reader) IntRange(max int) int {
	pBig, _ := crand.Int(m, big.NewInt(int64(max)))
	return int(pBig.Uint64()) + 1
}

func (m *reader) next() {
	m.vector = m.nextVector[:]
	next := sha256.Sum256(m.nextVector[:])
	m.nextVector = next[:]
	m.position = 0
}

func (m *reader) Read(b []byte) (n int, err error) {
	var lenB = len(b)
	var lenV = 32 - m.position

	if lenV < lenB {
		m.next()
	}

	n = copy(b, m.vector[m.position:m.position+lenB])
	m.position++

	return n, nil
}

func (m *reader) Shuffle(input string) string {
	b := []byte(input)
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	_, _ = h.Write([]byte(m.init))
	rand.Seed(int64(h.Sum64()))
	rand.Shuffle(len(input), func(i, j int) {
		b[i], b[j] = b[j], input[i]
	})

	return string(b)
}
