using System;
using System.Collections.Generic;

public static class DialingCodes
{
	private static Dictionary<int, string> emptyCodes = new Dictionary<int, string>();

	public static Dictionary<int, string> GetEmptyDictionary()
	{
		return emptyCodes;
	}

	public static Dictionary<int, string> GetExistingDictionary()
	{
		return new Dictionary<int, string>{
		{1, "United States of America"},
		{55, "Brazil"},
		{91, "India"}
	};
	}

	public static Dictionary<int, string> AddCountryToEmptyDictionary(int countryCode, string countryName)
	{
		var codes = new Dictionary<int, string>();
		codes.Add(countryCode, countryName);
		return codes;
	}

	public static Dictionary<int, string> AddCountryToExistingDictionary(
		Dictionary<int, string> existingDictionary, int countryCode, string countryName)
	{
		if (!existingDictionary.ContainsKey(countryCode))
		{
			existingDictionary.Add(countryCode, countryName);
		}
		return existingDictionary;
	}

	public static string GetCountryNameFromDictionary(
		Dictionary<int, string> existingDictionary, int countryCode)
	{
		if (!existingDictionary.ContainsKey(countryCode))
		{
			return string.Empty;
		}
		return existingDictionary[countryCode];
	}

	public static bool CheckCodeExists(Dictionary<int, string> existingDictionary, int countryCode)
	{
		return existingDictionary.ContainsKey(countryCode);
	}

	public static Dictionary<int, string> UpdateDictionary(
		Dictionary<int, string> existingDictionary, int countryCode, string countryName)
	{
		if (existingDictionary.ContainsKey(countryCode))
		{
			existingDictionary[countryCode] = countryName;
		}
		return existingDictionary;
	}

	public static Dictionary<int, string> RemoveCountryFromDictionary(
		Dictionary<int, string> existingDictionary, int countryCode)
	{
		if (existingDictionary.ContainsKey(countryCode))
		{
			existingDictionary.Remove(countryCode);
		}
		return existingDictionary;
	}

	public static string FindLongestCountryName(Dictionary<int, string> existingDictionary)
	{
		string result = "";
		foreach (var code in existingDictionary)
		{
			if (code.Value.Length > result.Length)
			{
				result = code.Value;
			}
		}
		return result;
	}
}