
#include "circular_buffer.h"

int16_t read(circular_buffer_t *buffer, buffer_value_t *read_value)
{
	if (buffer->length == 0)
	{
		errno = ENODATA;
		return (EXIT_FAILURE);
	}
	*read_value = buffer->values[buffer->start];
	buffer->start++;
	if (buffer->start == buffer->capacity)
	{
		buffer->start = 0;
	}
	buffer->length--;
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

void clear_buffer(circular_buffer_t *buffer)
{
	buffer->length = 0;
	buffer->start = 0;
	buffer->end = 0;
}

int16_t write_uni(circular_buffer_t *buffer, buffer_value_t value, bool writeover)
{
	if (buffer->length < buffer->capacity)
	{
		writeover = false;
	}
	if (writeover)
	{
		buffer->start++;
		if (buffer->start == buffer->capacity)
		{
			buffer->start = 0;
		}
	}
	else
	{
		if (buffer->length == buffer->capacity)
		{
			errno = ENOBUFS;
			return (EXIT_FAILURE);
		}
		buffer->length++;
	}
	buffer->values[buffer->end] = value;
	buffer->end++;
	if (buffer->end == buffer->capacity)
	{
		buffer->end = 0;
	}
	return (EXIT_SUCCESS);
}

int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value)
{
	return (write_uni(buffer, value, true));
}

int16_t write(circular_buffer_t *buffer, buffer_value_t value)
{
	return (write_uni(buffer, value, false));
}