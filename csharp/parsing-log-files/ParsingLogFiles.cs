using System;
using System.Text.RegularExpressions;

public class LogParser
{
    public bool IsValidLine(string text)
    {
        return Regex.IsMatch(text, @"^\[TRC\].*|^\[DBG\].*|^\[INF\].*|^\[WRN\].*|^\[ERR\].*|^\[FTL\].*");
    }

    public string[] SplitLogLine(string text)
    {
        string pattern = @"\<[\^\*\=\-]*\>";
        return Regex.Split(text, pattern);
    }

    public int CountQuotedPasswords(string lines)
    {
        int count = 0;
        string[] allLines = lines.Split(Environment.NewLine);
        foreach (string line in allLines)
        {
            if (Regex.IsMatch(line, "\\\".*password.*\\\"", RegexOptions.IgnoreCase))
            {
                count++;
            }
        }
        return count;
    }

    public string RemoveEndOfLineText(string line)
    {
        return Regex.Replace(line, @"end-of-line\d*", "");
    }

    public string[] ListLinesWithPasswords(string[] lines)
    {
        throw new NotImplementedException($"Please implement the LogParser.ListLinesWithPasswords() method");
    }
}
