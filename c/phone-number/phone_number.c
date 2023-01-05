#include "phone_number.h"


char *phone_number_clean(const char *input)
{
	char *result;
	char *result_copy;
	unsigned int i;

	result = malloc(sizeof(*result) * (NUMBER_LENGTH + 1));
	if (result == NULL)
	{
		return (NULL);
	}
	result_copy = result;
	i = 0;
	while (*input)
	{
		if (*input >= '0' && *input <= '9')
		{
			*result_copy = *input;
			result_copy++;
			i++;
		}
		input++;
	}
	if (i != NUMBER_LENGTH)
	{
		strcpy(result, "0000000000");
	}
	return (result);
}