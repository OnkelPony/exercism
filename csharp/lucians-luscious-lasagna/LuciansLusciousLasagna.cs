class Lasagna
{
	public const int TimeToCookLasagna = 40;
	public const int TimeToPrepareLayer = 2;
	// TODO: define the 'ExpectedMinutesInOven()' method
	public int ExpectedMinutesInOven()
	{
		return TimeToCookLasagna;
	}
	// TODO: define the 'RemainingMinutesInOven()' method
	public int RemainingMinutesInOven(int minutesCooking)
	{
		return TimeToCookLasagna - minutesCooking;
	}
	// TODO: define the 'PreparationTimeInMinutes()' method
	public int PreparationTimeInMinutes(int layers)
	{
		return layers * TimeToPrepareLayer;
	}
	// TODO: define the 'ElapsedTimeInMinutes()' method
    public int ElapsedTimeInMinutes(int layers, int minutesCooking)
    {
        return PreparationTimeInMinutes(layers) + minutesCooking;
    }
}
