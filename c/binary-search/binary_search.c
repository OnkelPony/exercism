#include "binary_search.h"

const int *binary_search(int value, const int *arr, size_t length)
{
	int i;
	int lower;
	int upper;

	if (length == 0)
	{
		return (NULL);
	}
	i = length / 2;
	lower = 0;
	upper = length - 1;
	while (1)
	{
		if (value > arr[i])
		{
			lower = i;
			i = upper - (upper - lower) / 2;
		}
		else if (value < arr[i])
		{
			upper = i;
			i = lower + (upper - lower) / 2;
		}
		else
		{
			return (&arr[i]);
		}
	}
	return (NULL);
}
