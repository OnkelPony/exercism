#include "rational_numbers.h"
#include <stdlib.h>

rational_t add(rational_t a, rational_t b)
{
	rational_t result;

	result.numerator = a.numerator * b.denominator + b.numerator * a.denominator;
	result.denominator = a.denominator * b.denominator;
	return (reduce(result));
}

rational_t subtract(rational_t a, rational_t b);
rational_t multiply(rational_t a, rational_t b);
rational_t divide(rational_t a, rational_t b);
rational_t absolute(rational_t r);
rational_t exp_rational(rational_t r, int16_t n);

rational_t reduce(rational_t r)
{
	rational_t result;
	int gcd;
	int d;

	gcd = 1;
	d = 1;
	while (d <= abs(r.numerator) && d <= abs(r.denominator))
	{
		if (r.numerator % d == 0 && r.denominator % d == 0)
		{
			gcd = d;
		}
		d++;
	}
	result.numerator = r.numerator / gcd;
	result.denominator = r.denominator / gcd;
	return (result);
}

float exp_real(uint16_t x, rational_t r);
