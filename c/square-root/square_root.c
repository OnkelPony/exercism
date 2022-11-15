#include "square_root.h"

unsigned int	square_root(unsigned int number)
{
	unsigned int	c;
	unsigned int	d;

	c = 0;
	d = 1 << 30;
	while (d > number)
		d >>= 2;
	while (d != 0) {
		if (number >= c + d)
		{
			number -= c + d;
			c = (c >> 1) + d;
		}
		else {
			c >>= 1;
		}
		d >>= 2;
	}
	return (c);
}