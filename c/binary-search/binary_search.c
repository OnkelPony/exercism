#include "binary_search.h"

const int *binary_search(int value, const int *arr, size_t length)
{
	if (arr[length - 1] == value)
	{
		return (&arr[length - 1]);
	}
	return (NULL);
}
