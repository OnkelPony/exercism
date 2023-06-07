using System;
using System.Collections.Generic;

public static class DialingCodes
{
	private static Dictionary<int, string> emptyCodes = new Dictionary<int, string>();
	private static Dictionary<int, string> smallCodes = new Dictionary<int, string>{
		{1, "United States of America"},
		{55, "Brasil"},
		{91, "India"}
	};

	public static Dictionary<int, string> GetEmptyDictionary()
	{
		return emptyCodes;
	}

	public static Dictionary<int, string> GetExistingDictionary()
	{
		return smallCodes;
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
		throw new NotImplementedException($"Please implement the (static) UpdateDictionary() method");
	}

	public static Dictionary<int, string> RemoveCountryFromDictionary(
		Dictionary<int, string> existingDictionary, int countryCode)
	{
		throw new NotImplementedException($"Please implement the (static) RemoveCountryFromDictionary() method");
	}

	public static string FindLongestCountryName(Dictionary<int, string> existingDictionary)
	{
		throw new NotImplementedException($"Please implement the (static) FindLongestCountryName() method");
	}
}