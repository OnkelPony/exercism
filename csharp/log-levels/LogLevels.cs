using System;

static class LogLine
{
	public static string Message(string logLine)
	{
		return logLine.Split(':')[1].Trim();
	}

	public static string LogLevel(string logLine)
	{
		return logLine.Split(':')[0].Trim('[', ']').ToLower();
	}

	public static string Reformat(string logLine)
	{
		string[] logElements = logLine.Split(':');
        return $"{logElements[1].Trim()} ({logElements[0].Trim('[', ']').ToLower()})";
	}
}
