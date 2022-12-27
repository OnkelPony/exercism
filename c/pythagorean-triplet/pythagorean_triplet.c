#include "pythagorean_triplet.h"

triplets_t *triplets_with_sum(uint16_t sum)
{
<<<<<<< HEAD
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
=======
	triplets_t *p_triplets;
	size_t a;
	size_t b;
	size_t c;
	size_t count;

	p_triplets = malloc(sizeof(*p_triplets));
	p_triplets->triplets = malloc(sizeof(p_triplets->triplets));
	p_triplets->count = 0;
	a = 1;
	c = sum;
	count = 0;
	while (a < sum / 3)
	{
		b = a + 1;
		while (b < c)
		{
			c = sum - a - b;
			if (c * c == a * a + b * b)
			{
				count++;
				p_triplets->triplets = realloc(p_triplets->triplets, sizeof((*p_triplets->triplets)) * count);
				p_triplets->triplets[count - 1].a = a;
				p_triplets->triplets[count - 1].b = b;
				p_triplets->triplets[count - 1].c = c;
				p_triplets->count = count;
			}
			b++;
		}
		a++;
	}
>>>>>>> edd6c77 (c_pythagorean-triplet complete.)
	return (p_triplets);
}

void free_triplets(triplets_t *triplets)
{
	free(triplets->triplets);
	free(triplets);
}
