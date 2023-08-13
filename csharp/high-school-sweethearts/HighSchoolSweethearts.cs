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
        return $@"
     ******       ******
   **      **   **      **
 **         ** **         **
**            *            **
**                         **
**     {studentA} +  {studentB}    **
 **                       **
   **                   **
     **               **
       **           **
         **       **
           **   **
             ***
              *
";
    }

    public static string DisplayGermanExchangeStudents(string studentA
        , string studentB, DateTime start, float hours)
    {
        var culture =     System.Globalization.CultureInfo.GetCultureInfo("de_DE");
        FormattableString message = $"{start} - that's {hours} hours";
        return $"{studentA} and {studentB} have been dating since {message.ToString(culture)}";
    }
}
