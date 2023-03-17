package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	switch order {
	case "burger":
		{
			food := food{preptime: 15}
			return food.preptime
		}
	case "nuggets":
		{
			food := food{preptime: 12}
			return food.preptime
		}
	case "chips":
		{
			food := food{preptime: 10}
			return food.preptime
		}
	}
	return 404
}
