package allergies

const testVersion = 1

// Allergen ..
type Allergen struct {
	val  uint
	name string
}

var allergenList = []Allergen{
	{1, "eggs"},
	{1 << 1, "peanuts"},
	{1 << 2, "shellfish"},
	{1 << 3, "strawberries"},
	{1 << 4, "tomatoes"},
	{1 << 5, "chocolate"},
	{1 << 6, "pollen"},
	{1 << 7, "cats"},
}

// Allergies ..
func Allergies(input uint) []string {
	var affectedBy []string
	for _, all := range allergenList {
		if input&all.val == all.val {
			affectedBy = append(affectedBy, all.name)
		}
	}
	return affectedBy
}

// AllergicTo ..
func AllergicTo(input uint, allergen string) bool {
	for _, all := range allergenList {
		if all.name == allergen {
			return input&all.val == all.val
		}
	}

	return false
}
