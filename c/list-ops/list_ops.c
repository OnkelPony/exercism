#include "list_ops.h"

list_t *new_list(size_t length, list_element_t elements[])
{
	list_t *p_list;

	p_list = malloc(sizeof(*p_list));
	// p_list->elements = elements;
	// p_list->length = length;
	if (elements)
	{
		elements[0] = length;
	}
	return (p_list);
}

void delete_list(list_t *list)
{
	free(list);
}

list_t *append_list(list_t *list1, list_t *list2)
{
	list_t *p_list;
	// size_t length;

	p_list = malloc(sizeof(*p_list));
	list1->length = list2->length;
	// p_list->length = length;
	return (p_list);
}
