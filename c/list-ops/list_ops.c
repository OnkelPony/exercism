#include "list_ops.h"

list_t *new_list(size_t length, list_element_t elements[])
{
	list_t *p_list;
	unsigned int i;

	i = 0;
	p_list = malloc(sizeof(*p_list) + length * sizeof(p_list->elements[0]));
	p_list->length = length;
	while (i < length)
	{
		p_list->elements[i] = elements[i];
		i++;
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
	unsigned int i;

	i = 0;
	p_list = new_list(list1->length, list1->elements);
	p_list = realloc(p_list, sizeof(*p_list) + (list1->length + list2->length) * sizeof(p_list->elements[0]));
	p_list->length = list1->length + list2->length;
	while (i < list2->length)
	{
		p_list->elements[i + list1->length] = list2->elements[i];
		i++;
	}
	return (p_list);
}

list_t *filter_list(list_t *list, bool (*filter)(list_element_t))
{
	list_t *p_list;
	unsigned int i;

	i = 0;
	p_list = malloc(sizeof(*p_list) + list->length * sizeof(p_list->elements[0]));
	p_list->length = 0;
	while (i < list->length)
	{
		if (filter((int)list->elements[i]))
		{
			p_list->elements[p_list->length] = list->elements[i];
			p_list->length++;
		}
		i++;
	}
	return (p_list);
}

list_t *map_list(list_t *list, list_element_t (*map)(list_element_t))
{
	list_t *p_list;
	unsigned int i;

	i = 0;
	p_list = malloc(sizeof(*p_list) + list->length * sizeof(p_list->elements[0]));
	p_list->length = list->length;
	while (i < list->length)
	{
		p_list->elements[i] = map(list->elements[i]);
		i++;
	}
	return (p_list);
}

size_t length_list(list_t *list)
{
	return (list->length);
}

list_element_t foldl_list(list_t *list, list_element_t initial,
						  list_element_t (*foldl)(list_element_t,
												  list_element_t))
{
	list_element_t result;
	unsigned int i;

	result = initial;
	i = 0;
	while (i < list->length)
	{
		result = foldl(list->elements[i], initial);
		i++;
	}
	return (result);
}