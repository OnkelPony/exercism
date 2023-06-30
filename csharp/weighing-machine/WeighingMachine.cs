using System;

class WeighingMachine
{
    private double _weight;
    public WeighingMachine(int precision)
    {
        Precision = precision;
    }
    public int Precision { get; set; }
    public double Weight
    {
        get => _weight;
        set
        {
            if (value < 0)
            {
                throw new ArgumentOutOfRangeException();
            }
            _weight = value;
        }
    }
    public string DisplayWeight => String.Format("{0:N" + Precision + "} kg", _weight - TareAdjustment);
    public double TareAdjustment { get; set; } = 5;
}
