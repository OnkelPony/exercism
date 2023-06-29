using System;

public static class CentralBank
{
	public static string DisplayDenomination(long @base, long multiplier)
	{
		long result;
		try
		{
			result = checked(@base * multiplier);
		}
		catch (OverflowException)
		{
			return "*** Too Big ***";
		}
		return $"{result}";
	}

	public static string DisplayGDP(float @base, float multiplier)
	{
		float result = @base * multiplier;
		if (Double.IsInfinity(result))
		{
			return "*** Too Big ***";
		}
		return $"{result}";
	}

	public static string DisplayChiefEconomistSalary(decimal salaryBase, decimal multiplier)
	{
		decimal result;
		try
		{
			result = checked(salaryBase * multiplier);
		}
		catch (OverflowException)
		{
			return "*** Much Too Big ***";
		}
		return $"{result}";
	}
}
