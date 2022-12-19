#include "list_ops.h"

list_t *new_list(size_t length, list_element_t elements[])
{
	list_t *p_list;

	p_list = malloc(sizeof(*p_list) * length);
	p_list->elements = elements;
	p_list->length = length;
	return (p_list);
}

void delete_list(list_t *list)
{
	free(list);
}
