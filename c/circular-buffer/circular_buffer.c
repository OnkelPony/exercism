#include <stdlib.h>
#include "circular_buffer.h"

int16_t read(circular_buffer_t *buffer, buffer_value_t *read_value);

circular_buffer_t *new_circular_buffer(size_t capacity)
{
	circular_buffer_t *p_buffer;

	p_buffer = malloc(sizeof(*p_buffer) + capacity * sizeof(p_buffer->values[0]));
	return (p_buffer);
}

void delete_buffer(circular_buffer_t *buffer)
{
	free(buffer);
}

void clear_buffer(circular_buffer_t *buffer);

int16_t write(circular_buffer_t *buffer, buffer_value_t value);

int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value)
{
	buffer->end = value;
	return(EXIT_SUCCESS);
}