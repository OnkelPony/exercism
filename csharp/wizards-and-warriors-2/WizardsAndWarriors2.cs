using System;

static class GameMaster
{
	public static string Describe(Character character)
	{
		return $"You're a level {character.Level} {character.Class} with {character.HitPoints} hit points.";
	}

	public static string Describe(Destination destination)
	{
        return $"You've arrived at {destination.Name}, which has {destination.Inhabitants} inhabitants.";
	}

	public static string Describe(TravelMethod travelMethod)
	{
        string method = travelMethod switch
        {
            TravelMethod.Walking => "by walking",
            TravelMethod.Horseback => "on horseback",
            _ => "whatever"
        };
        return $"You're traveling to your destination {method}.";
	}

	public static string Describe(Character character, Destination destination, TravelMethod travelMethod)
	{
		throw new NotImplementedException("Please implement the (static) GameMaster.Describe(Character, Destination, TravelMethod) method");
	}

	public static string Describe(Character character, Destination destination)
	{
		throw new NotImplementedException("Please implement the (static) GameMaster.Describe(Character, Destination) method");
	}
}

class Character
{
	public string Class { get; set; }
	public int Level { get; set; }
	public int HitPoints { get; set; }
}

class Destination
{
	public string Name { get; set; }
	public int Inhabitants { get; set; }
}

enum TravelMethod
{
	Walking,
	Horseback
}
