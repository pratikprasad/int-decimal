package intdecimal

import (
	"fmt"
	"math/big"
)

// Stores a decimal as a big integer shifted 10^6
// Gives us 6 degrees of precision and arbitrary range
type Decimal struct{ i big.Int }

const microMultiplier = 1000000

var bigMicroMultiplier = big.NewInt(1000000)

func NewDecimal(f float64) Decimal {
	return Decimal{i: *big.NewInt(int64(f * microMultiplier))}
}

func (d Decimal) Add(d2 Decimal) Decimal {
	_ = *d.i.Add(&d.i, &d2.i)
	return d
}

//func (d Decimal) Div(d2 Decimal) Decimal {
//	return d / d2
//}
//
//func (d Decimal) Mul(d2 Decimal) Decimal {
//	return d * d2
//}
//
//func (d Decimal) Sub(d2 Decimal) Decimal {
//	return d - d2
//}

func (d Decimal) String() string {
	intPart := big.NewInt(0).Div(&d.i, bigMicroMultiplier).Int64()
	decPart := big.NewInt(0).Mod(&d.i, bigMicroMultiplier).Int64()
	return fmt.Sprintf("%d.%d", intPart, decPart)
}
