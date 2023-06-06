using System;

public static class Identifier
{
	public static string Clean(string identifier)
	{
		string result = "";
		bool toCamel = false;
		foreach (char ch in identifier)
		{
			if (ch == ' ')
			{
				result += '_';
			}
			else if (ch == '-')
			{
				toCamel = true;
			}
			else if (ch == '\0')
			{
				result += "CTRL";
			}
			else if (!Char.IsLetter(ch) || (ch >= '\u03B1' && ch <= '\u03C9'))
			{
			}
			else
			{
				if (toCamel)
				{
					result += Char.ToUpperInvariant(ch);
					toCamel = false;
				}
				else
				{
					result += ch;
				}
			}
		}
		return result;
	}
}
