#include "grade_school.h"

void init_roster(roster_t *roster)
{
	roster->count = 0;
}

bool add_student(roster_t *roster, char *name, int grade)
{
	bool succeed;

	roster->count = 0;
	strcpy(roster.students[0].name, name);
	succeed = true;
	return succeed;
}

roster_t get_grade(roster_t *roster, int grade)
{
	roster_t result;

	return result;
}