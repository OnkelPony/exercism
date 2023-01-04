#include "phone_number.h"


char *phone_number_clean(const char *input)
{
	char *result;
	char *result_copy;

	result = malloc(sizeof(*result) * (NUMBER_LENGTH + 1));
	result_copy = result;
	while (*input)
	{
		if (*input >= '0' && *input <= '9')
		{
			*result_copy = *input;
			result_copy++;
		}
		input++;
	}
	return (result);
}