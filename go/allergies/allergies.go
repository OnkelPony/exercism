package allergies

var allergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

// Allergies returns slice of allergies based on score in parameter.
func Allergies(score uint) []string {
	score = score % 256
	var result []string
	for i, allergen := range allergens {
		if 1<<i == 1<<i & score{
			result = append(result, allergen)
		}
	}
	return result
}

// AllergicTo returns true if score contains substance entered as parameters.
func AllergicTo(score uint, substance string) bool {
	result := Allergies(score)
	for _, allergen := range result {
		if allergen == substance {
			return true
		}
	}
	return false
}
