using System;

class WeighingMachine
{
	private int _precision;
	public WeighingMachine(int precision)
	{
		_precision = precision;
	}
	public int Precision => _precision;
	public float Weight { get; private set; }
	// TODO: define the 'DisplayWeight' property
    public string DisplayWeight => $"{Weight} kg";
	public float TareAdjustment { get; private set; }
}
