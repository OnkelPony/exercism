#include "allergies.h"

allergen_list_t get_allergens(unsigned int score)
{
	allergen_list_t allergen_list;
	unsigned int i;
	unsigned int remainder;

	allergen_list.count = 0;
	score %= 256;
	remainder = score;
	i = ALLERGEN_COUNT - 1;
	while (remainder > 0)
	{
		remainder = score % (unsigned int)pow(2, i);
		if (remainder != score)
		{
			score = remainder;
			allergen_list.allergens[i] = true;
			allergen_list.count++;
		}
		i--;
	}
	return (allergen_list);
}

bool is_allergic_to(allergen_t allergen, unsigned int score)
{
	allergen_list_t allergen_list;

	allergen_list = get_allergens(score);
	return (allergen_list.allergens[allergen]);
}
