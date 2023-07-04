using System;
using System.Collections.Generic;

// TODO implement the IRemoteControlCar interface
public interface IRemoteControlCar
{
    public int DistanceTravelled { get; }
	void Drive();
}

public class ProductionRemoteControlCar : IRemoteControlCar, IComparable<T>
{
	public int DistanceTravelled { get; private set; }
	public int NumberOfVictories { get; set; }

	public void Drive()
	{
		DistanceTravelled += 10;
	}

	int IComparable<T>.CompareTo(T? other)
	{
		throw new NotImplementedException();
	}
}

public class ExperimentalRemoteControlCar : IRemoteControlCar
{
	public int DistanceTravelled { get; private set; }

	public void Drive()
	{
		DistanceTravelled += 20;
	}
}

public static class TestTrack
{
	public static void Race(IRemoteControlCar car)
	{
        car.Drive();
	}

	public static List<ProductionRemoteControlCar> GetRankedCars(ProductionRemoteControlCar prc1,
		ProductionRemoteControlCar prc2)
	{
		throw new NotImplementedException($"Please implement the (static) TestTrack.GetRankedCars() method");
	}
}
