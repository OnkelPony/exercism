#ifndef CIRCULAR_BUFFER_H
#define CIRCULAR_BUFFER_H
#include <stdint.h>
#include <stddef.h>
#define ARRAY_LENGTH(A) (sizeof(A) / sizeof(A[0]))

typedef int buffer_value_t;
typedef struct circular_buffer_t
{
	size_t capacity;
	size_t length;
	size_t start;
	size_t end;
	buffer_value_t values[];

} circular_buffer_t;
int16_t read(circular_buffer_t *buffer, buffer_value_t *read_value);
circular_buffer_t *new_circular_buffer(size_t capacity);
void delete_buffer(circular_buffer_t *buffer);
void clear_buffer(circular_buffer_t *buffer);
int16_t write(circular_buffer_t *buffer, buffer_value_t value);
int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value);
#endif
