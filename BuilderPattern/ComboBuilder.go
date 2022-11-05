package BuilderPattern

// Friendship Day Bucket contains two Hot Wings and a Medium Fry
func BuildFriendshipDayBucket(coupon *Coupon) Combo {
	combo := Combo{coupon: coupon, name: "Friendship Day Bucket"}

	combo.addItem(HotWings{})
	combo.addItem(HotWings{})
	combo.addItem(MediumFries{})

	return combo
}

// Vegan Special Bucket contains two Medium Fries
func BuildVeganSpecialBucket(coupon *Coupon) Combo {
	combo := Combo{coupon: coupon, name: "Vegan Special Bucket"}

	combo.addItem(MediumFries{})
	combo.addItem(MediumFries{})

	return combo
}
