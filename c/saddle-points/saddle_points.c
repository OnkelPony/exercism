#include "saddle_points.h"

void free_saddle_points(saddle_points_t *saddle_points)
{
	free(saddle_points);
}

saddle_points_t *saddle_points(uint8_t rows, uint8_t columns, uint8_t matrix[rows][columns])
{
	saddle_points_t *row_max;
	saddle_points_t *col_min;
	saddle_points_t *result;

	row_max = get_row_max(rows, columns, matrix);
	col_min = get_col_min(rows, columns, matrix);
	result = intersec(row_max, col_min);
	free(row_max);
	free(col_min);
	return (result);
}

saddle_points_t *get_row_max(uint8_t rows, uint8_t columns, uint8_t matrix[rows][columns])
{
	saddle_points_t *row_max;
	uint8_t max;
	uint8_t i;
	uint8_t j;

	i = 0;
	row_max = malloc(sizeof(*row_max));
	row_max->count = 0;
	row_max->points = malloc(sizeof(*row_max->points) * rows * columns);
	while (i < rows)
	{
		j = 0;
		max = 0;
		while (j < columns)
		{
			if (matrix[i][j] > max)
			{
				max = matrix[i][j];
			}
			j++;
		}
		j = 0;
		while (j < columns)
		{
			if (matrix[i][j] == max)
			{
				row_max->points[row_max->count].row = i + 1;
				row_max->points[row_max->count].column = j + 1;
				row_max->count++;
			}
			j++;
		}
		i++;
	}
	return (row_max);
}

saddle_points_t *get_col_min(uint8_t rows, uint8_t columns, uint8_t matrix[rows][columns])
{
	saddle_points_t *col_min;
	uint8_t min;
	uint8_t i;
	uint8_t j;

	i = 0;
	col_min = malloc(sizeof(*col_min));
	col_min->count = 0;
	col_min->points = malloc(sizeof(*col_min->points) * rows * columns);
	while (i < columns)
	{
		j = 0;
		min = 255;
		while (j < rows)
		{
			if (matrix[j][i] < min)
			{
				min = matrix[j][i];
			}
			j++;
		}
		j = 0;
		while (j < rows)
		{
			if (matrix[j][i] == min)
			{
				col_min->points[col_min->count].row = j + 1;
				col_min->points[col_min->count].column = i + 1;
				col_min->count++;
			}
			j++;
		}
		i++;
	}
	return (col_min);
}

saddle_points_t *intersec(const saddle_points_t *row_max, const saddle_points_t *col_min)
{
	saddle_points_t *result;
	size_t i;
	size_t j;

	result = malloc(sizeof(*result));
	result->count = 0;
	result->points = malloc(sizeof(*result->points) * row_max->count * col_min->count);
	i = 0;
	while (i < row_max->count)
	{
		j = 0;
		while (j < col_min->count)
		{
			if (row_max->points[i].row == col_min->points[j].row && row_max->points[i].column == col_min->points[j].column)
			{
				result->points[result->count] = row_max->points[i];
				result->count++;
			}
			j++;
		}
		i++;
	}
	return (result);
}