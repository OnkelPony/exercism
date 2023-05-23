using System;

public static class LogAnalysis
{
    // TODO: define the 'SubstringAfter()' extension method on the `string` type
public static string SubstringAfter(this string str, string separator)
{
    int i = str.IndexOf(separator);
    return i != -1? str.Substring(i + separator.Length) : str;
}
    // TODO: define the 'SubstringBetween()' extension method on the `string` type

    // TODO: define the 'Message()' extension method on the `string` type

    // TODO: define the 'LogLevel()' extension method on the `string` type
}