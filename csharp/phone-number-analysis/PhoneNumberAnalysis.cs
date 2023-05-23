using System;

public static class PhoneNumber
{
	public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber)
	{
		string[] numberParts = phoneNumber.Split('-');
		return (numberParts[0] == "212", numberParts[1] == "555", numberParts[2]);
	}

	public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo)
	{
		return phoneNumberInfo.IsFake;
	}
}
