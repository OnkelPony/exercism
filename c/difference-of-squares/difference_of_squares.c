#include "difference_of_squares.h"

unsigned int sum_of_squares(unsigned int number)
{
	unsigned int    result;
	unsigned int	i;

	result = 0;
	i = 0;
	while (i <= number)
	{
		result += i * i;
		i++;
	}
	return (result);
}

unsigned int square_of_sum(unsigned int number)
{
	unsigned int    result;

    result = number;
    return (result);
}

unsigned int difference_of_squares(unsigned int number)
{
	unsigned int    result;

    result = number;
    return (result);
}