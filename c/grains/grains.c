#include "grains.h"

uint64_t square(uint8_t index)
{
    uint64_t    result;

    result = 1;
    if (index < 1 || index > 64)
    {
        return (0);
    }
    while (index > 1)
    {
        result *= 2;
        index--;
    }
    return (result);
}
uint64_t total(void)
{
    uint64_t    result;
    int         i;

    result = 0;
    i = 1;
    while (i <= 64)
    {
        result += square(i);
        i++;
    }
    return (result);
}