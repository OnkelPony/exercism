#ifndef PYTHAGOREAN_TRIPLET_H
#define PYTHAGOREAN_TRIPLET_H
#include <stdint.h>
#include <stdlib.h>
typedef struct triplet_t
{
	size_t a;
	size_t b;
	size_t c;
} triplet_t;

typedef struct triplets_t
{
	size_t count;
	triplet_t *triplets;
} triplets_t;

triplets_t *triplets_with_sum(uint16_t sum);
void free_triplets(triplets_t *triplets);

#endif
