package set

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && (1<<bit)&s.words[word] != 0
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for word, bits := range s.words {
		if bits == 0 {
			continue
		}
		for bit := 0; bit < 64; bit++ {
			if (1<<uint(bit))&bits != 0 {
				if buf.Len() > len("[") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*word+bit)
			}
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
