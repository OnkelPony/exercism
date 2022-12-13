#include "grade_school.h"

void init_roster(roster_t *roster)
{
	roster->count = 0;
}

bool add_student(roster_t *roster, char *name, int grade)
{
	bool succeed;

	strcpy(roster->students[roster->count].name, name);
	roster->students[roster->count].grade = grade;
	(roster->count)++;
	succeed = true;
	return succeed;
}

roster_t get_grade(roster_t *roster, int grade)
{
	roster_t result;

	result = *roster;
	result.students[0].grade = grade;
	return result;
}