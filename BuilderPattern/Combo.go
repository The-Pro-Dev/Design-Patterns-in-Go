package BuilderPattern

type Combo struct {
	name   string
	items  []FoodItem
	coupon *Coupon
}

func (combo *Combo) addItem(item FoodItem) {
	(*combo).items = append(combo.items, item)
}

func (combo Combo) getItems() []string {
	itemCount := len(combo.items)

	items := make([]string, itemCount)

	for i := 0; i < itemCount; i++ {
		items[i] = combo.items[i].getName()
	}

	return items
}

func (combo Combo) getName() string {
	return combo.name
}

func (combo Combo) getPrice() float64 {
	price := 0.0

	itemCount := len(combo.items)

	for i := 0; i < itemCount; i++ {
		price += combo.items[i].getPrice()
	}

	if combo.coupon != nil {
		coupon := *combo.coupon
		var err error
		price, err = coupon.applyDiscount(price)

		if err == nil {
			return price
		}
	}

	return price
}
