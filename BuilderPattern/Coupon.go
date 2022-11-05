package BuilderPattern

import "errors"

type CouponKind string

const (
	FlatDiscount       = CouponKind("Flat")
	PercentageDiscount = CouponKind("Percentage")
)

type Coupon struct {
	code     string
	kind     CouponKind
	discount float64
	minValue float64
}

func (coupon Coupon) applyDiscount(price float64) (float64, error) {
	if price < coupon.minValue {
		return price, nil
	}

	if coupon.kind == FlatDiscount {
		return price - coupon.discount, nil
	}

	if coupon.kind == PercentageDiscount {
		return price * (1 - coupon.discount), nil
	}

	return -1, errors.New("invalid coupon kind")
}
