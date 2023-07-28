using System;

public static class HighSchoolSweethearts
{
    public static string DisplaySingleLine(string studentA, string studentB)
    {
        const int width = 61;
        string result = $"{studentA} ♡ {studentB}";
        int leftPadding = (width - result.Length) / 2 - 1;
        int rightPadding = (width - leftPadding - result.Length);
        return $"{new string(' ', leftPadding)}{result}{new string(' ', rightPadding)}";
    }

    public static string DisplayBanner(string studentA, string studentB)
    {
        throw new NotImplementedException($"Please implement the (static) HighSchoolSweethearts.DisplayBanner() method");
    }

    public static string DisplayGermanExchangeStudents(string studentA
        , string studentB, DateTime start, float hours)
    {
        throw new NotImplementedException($"Please implement the (static) HighSchoolSweethearts.DisplayGermanExchangeStudents() method");
    }
}
