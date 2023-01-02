#ifndef SADDLE_POINTS_H
#define SADDLE_POINTS_H
#include <stdint.h>
#include <stdlib.h>

typedef struct saddle_point_t
{
	uint8_t row;
	uint8_t column;
} saddle_point_t;

typedef struct saddle_points_t
{
	size_t count;
	saddle_point_t *points;
} saddle_points_t;

saddle_points_t *intersec(const saddle_points_t *row_max, const saddle_points_t *col_min);
saddle_points_t *get_row_max(uint8_t rows, uint8_t columns, uint8_t matrix[rows][columns]);
saddle_points_t *get_col_min(uint8_t rows, uint8_t columns, uint8_t matrix[rows][columns]);
saddle_points_t *saddle_points(uint8_t rows, uint8_t columns, uint8_t matrix[rows][columns]);
void free_saddle_points(saddle_points_t *saddle_points);
#endif
