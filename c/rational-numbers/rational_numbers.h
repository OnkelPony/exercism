#ifndef RATIONAL_NUMBERS_H
#define RATIONAL_NUMBERS_H

#include <stdint.h>
typedef struct rational_t
{
	int numerator;
	int denominator;
} rational_t;

rational_t add(rational_t a, rational_t b);
rational_t subtract(rational_t a, rational_t b);
rational_t multiply(rational_t a, rational_t b);
rational_t divide(rational_t a, rational_t b);
rational_t absolute(rational_t r);
rational_t exp_rational(rational_t r, int16_t n);
rational_t reduce(rational_t r);
float exp_real(uint16_t x, rational_t r);

#endif
