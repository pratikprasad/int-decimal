package intdecimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimal_Add(t *testing.T) {
	d := NewDecimal(0)
	for i := 0; i < 1000; i++ {
		d = d.Add(NewDecimal(0.01))
	}
	assert.Equal(t, "10.0", d.String())
}

func TestDecimal_Sub(t *testing.T) {
	//d := NewDecimal(10.0)
	//for i := 0; i < 1000; i++ {
	//	d = d.Sub(NewDecimal(0.01))
	//}
	//assert.Equal(t, 0.0, d.Float())
}
