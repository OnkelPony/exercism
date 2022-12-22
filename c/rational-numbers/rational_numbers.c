#include "rational_numbers.h"
#include <stdlib.h>
#include <math.h>

rational_t add(rational_t a, rational_t b)
{
	rational_t result;

	result.numerator = a.numerator * b.denominator + b.numerator * a.denominator;
	result.denominator = a.denominator * b.denominator;
	return (reduce(result));
}

rational_t subtract(rational_t a, rational_t b)
{
	b.numerator *= -1;
	return (add(a, b));
}

rational_t multiply(rational_t a, rational_t b)
{
	rational_t result;

	result.numerator = a.numerator * b.numerator;
	result.denominator = a.denominator * b.denominator;
	return (reduce(result));
}

rational_t divide(rational_t a, rational_t b)
{
	rational_t result;

	result.numerator = a.numerator * b.denominator;
	result.denominator = a.denominator * b.numerator;
	return (reduce(result));
}

rational_t absolute(rational_t r)
{
	r.numerator = abs(r.numerator);
	r.denominator = abs(r.denominator);
	return (reduce(r));
}

rational_t exp_rational(rational_t r, int16_t n)
{
	rational_t result;
	if (n < 0)
	{
		n = abs(n);
		result.numerator = pow(r.denominator, n);
		result.denominator = pow(r.numerator, n);
	}
	else
	{
		result.numerator = pow(r.numerator, n);
		result.denominator = pow(r.denominator, n);
	}
	return (reduce(result));
}

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
	if (result.numerator == 0)
	{
		result.denominator = 1;
	}
	if (result.denominator < 0)
	{
		result.numerator *= -1;
		result.denominator *= -1;
	}
	return (result);
}

float exp_real(uint16_t x, rational_t r)
{
	return (pow(x, (float)r.numerator / r.denominator));
}
