using System;

class BirdCount
{
    private int[] birdsPerDay;
    private static int[] birdsLastWeek = { 0, 2, 5, 3, 7, 8, 4 };

    public BirdCount(int[] birdsPerDay)
    {
        this.birdsPerDay = birdsPerDay;
    }

    public static int[] LastWeek()
    {
        return birdsLastWeek;
    }

    public int Today()
    {
        return this.birdsPerDay[birdsPerDay.Length - 1];
    }

    public void IncrementTodaysCount()
    {
        this.birdsPerDay[birdsPerDay.Length - 1]++;
    }

    public bool HasDayWithoutBirds()
    {
        foreach (int bird in this.birdsPerDay)
        {
            if (bird == 0)
            {
                return true;
            }
        }
        return false;
    }

    public int CountForFirstDays(int numberOfDays)
    {
        int total = 0;
        for (int day = 0; day < numberOfDays; day++)
        {
            total += this.birdsPerDay[day];
        }
        return total;
    }

    public int BusyDays()
    {
        int busyDays = 0;
        foreach (int bird in this.birdsPerDay)
        {
            if (bird >= 5)
            {
                busyDays++;
            }
        }
        return busyDays;
    }
}
