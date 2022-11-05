package BuilderPattern

import (
	"fmt"
	"strings"
)

func display(combo Combo) {
	fmt.Println(combo.getName(), "contains (", strings.Join(combo.getItems(), " + "), ") and costs", combo.getPrice())
}

func Main() {
	fmt.Println("** Builder Pattern **")

	coupon10pctOff := Coupon{code: "Hot Wings Special (10% off over ₹100)", minValue: 100.00, discount: 0.10, kind: PercentageDiscount}
	friendShipDayBucket := BuildFriendshipDayBucket(&coupon10pctOff)
	display(friendShipDayBucket)

	veganSpecialBucket := BuildVeganSpecialBucket(nil)
	display(veganSpecialBucket)
}
