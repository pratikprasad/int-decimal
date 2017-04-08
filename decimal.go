package intdecimal

import (
	"fmt"
	"math/big"
)

// Stores a decimal as a big integer shifted 10^6
// Gives us 6 degrees of precision and arbitrary range
type Decimal struct{ i big.Int }

const microMultiplier = 1000000

// Decimal of 1 is represented underneath as 10**6
var decMultiplier = NewDecimal(1)

func NewDecimal(f float64) Decimal {
	return Decimal{i: *big.NewInt(int64(f * microMultiplier))}
}

func (d Decimal) Add(d2 Decimal) Decimal {
	out := Decimal{}
	_ = out.i.Add(&d.i, &d2.i)
	return out
}

func (d Decimal) Sub(d2 Decimal) Decimal {
	out := Decimal{}
	_ = out.i.Sub(&d.i, &d2.i)
	return out
}

func (d Decimal) Div(d2 Decimal) Decimal {
	out := Decimal{}
	_ = out.i.Div(&d.i, &d2.i)
	return out
}

func (d Decimal) Mul(d2 Decimal) Decimal {
	out := Decimal{}
	_ = out.i.Mul(&d.i, &d2.i)
	return out
}

func (d Decimal) Mod(d2 Decimal) Decimal {
	out := Decimal{}
	_ = out.i.Mod(&d.i, &d2.i)
	return out
}

func (d Decimal) int64() int64 {
	return d.i.Int64()
}

func (d Decimal) String() string {
	intPart := d.Div(decMultiplier).int64()
	decPart := d.Mod(decMultiplier).int64()
	return fmt.Sprintf("%d.%d", intPart, decPart)
}
