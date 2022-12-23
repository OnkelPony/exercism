#include <stdlib.h>
#include <errno.h>
#include "circular_buffer.h"

int16_t read(circular_buffer_t *buffer, buffer_value_t *read_value)
{
	*read_value = buffer->values[buffer->start];
	if (buffer->length == 0)
	{
		errno = ENODATA;
		return (EXIT_FAILURE);
	}
	return (EXIT_SUCCESS);
}

circular_buffer_t *new_circular_buffer(size_t capacity)
{
	circular_buffer_t *p_buffer;

	p_buffer = malloc(sizeof(*p_buffer) + capacity * sizeof(p_buffer->values[0]));
	p_buffer->capacity = capacity;
	p_buffer->start = 0;
	p_buffer->end = 0;
	p_buffer->length = 0;
	return (p_buffer);
}

void delete_buffer(circular_buffer_t *buffer)
{
	free(buffer);
}

void clear_buffer(circular_buffer_t *buffer);

int16_t write(circular_buffer_t *buffer, buffer_value_t value)
{
	buffer->values[buffer->end] = value;
	buffer->end++;
	buffer->length++;
	return (EXIT_SUCCESS);
}

int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value)
{
	buffer->values[buffer->end] = value;
	return (EXIT_SUCCESS);
}