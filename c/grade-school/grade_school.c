#include <string.h>
#include <stdlib.h>
#include "grade_school.h"

int cmpfunc(const void *a, const void *b);

int cmpfunc(const void *a, const void *b)
{
	const student_t *p_a = a;
	const student_t *p_b = b;
	int grade_a = p_a->grade;
	int grade_b = p_b->grade;
	const char *p_name_a = p_a->name;
	const char *p_name_b = p_b->name;
	if (grade_a == grade_b)
	{
		return (strcmp(p_name_a, p_name_b));
	}
	return (grade_a - grade_b);
}

void init_roster(roster_t *roster)
{
	roster->count = 0;
}

bool add_student(roster_t *roster, const char *name, int grade)
{
	unsigned int i;

	i = 0;
	while (i < roster->count)
	{
		if (!strcmp(roster->students[i].name, name))
		{
			return (false);
		}
		i++;
	}
	strcpy(roster->students[roster->count].name, name);
	roster->students[roster->count].grade = grade;
	(roster->count)++;
	qsort(roster->students, roster->count, sizeof(student_t), cmpfunc);
	return (true);
}

roster_t get_grade(const roster_t *roster, int grade)
{
	roster_t result;
	unsigned int i;

	init_roster(&result);
	i = 0;
	while (i < roster->count)
	{
		if (roster->students[i].grade == grade)
		{
			result.students[result.count] = roster->students[i];
			result.count++;
		}
		i++;
	}
	return result;
}