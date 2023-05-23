using System;

public static class LogAnalysis
{
	// TODO: define the 'SubstringAfter()' extension method on the `string` type
	public static string SubstringAfter(this string str, string separator)
	{
		int i = str.IndexOf(separator);
		return i != -1 ? str.Substring(i + separator.Length) : str;
	}
	// TODO: define the 'SubstringBetween()' extension method on the `string` type
	public static string SubstringBetween(this string str, string before, string after)
	{
		int i = str.IndexOf(before);
		int j = str.IndexOf(after);
		return i != -1 && j != -1 ? str.Substring(i + before.Length, j - i - before.Length) : str;
	}
	// TODO: define the 'Message()' extension method on the `string` type
    public static string Message(this string str)
    {
        return str.Split(':')[1].Trim();
    }
	// TODO: define the 'LogLevel()' extension method on the `string` type
    public static string LogLevel(this string str)
    {
        return str.SubstringBetween("[", "]");
    }
}