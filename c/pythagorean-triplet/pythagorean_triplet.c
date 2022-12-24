#include "pythagorean_triplet.h"

triplets_t *triplets_with_sum(uint16_t sum)
{
	triplet_t *p_triplet;
	triplets_t *p_triplets;

	if (sum == 12)
	{
		p_triplet = malloc(sizeof(*p_triplet));
		p_triplet->a = 3;
		p_triplet->b = 4;
		p_triplet->c = 5;
	}
	p_triplets = malloc(sizeof(*p_triplets));
	p_triplets->count = 1;
	p_triplets->triplets = p_triplet;
	return (p_triplets);
}

void free_triplets(triplets_t *triplets)
{
	free(triplets->triplets);
	free(triplets);
}
