using System;

static class Badge
{
	public static string Print(int? id, string name, string? department)
	{
		if (department == null)
		{
			department = "OWNER";
		}
		if (id == null)
		{
					return $"{name} - {department.ToUpper()}";
		}
		return $"[{id}] - {name} - {department.ToUpper()}";
	}
}
