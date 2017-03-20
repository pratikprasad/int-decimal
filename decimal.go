package intdecimal

type Decimal int64

const microMultiplier = 1000000

func NewDecimal(f float64) Decimal {
	return Decimal(f * microMultiplier)
}

func (d Decimal) Add(d2 Decimal) Decimal {
	return d + d2
}

func (d Decimal) Div(d2 Decimal) Decimal {
	return d / d2
}

func (d Decimal) Mul(d2 Decimal) Decimal {
	return d * d2
}

func (d Decimal) Sub(d2 Decimal) Decimal {
	return d - d2
}

func (d Decimal) Float() float64 {
	return float64(d)/microMultiplier
}
