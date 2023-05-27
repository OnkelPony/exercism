using System;

class RemoteControlCar
{
    private int _distanceDriven = 0;
    private int _batteryCharge = 100;
    public static RemoteControlCar Buy()
    {
        return new RemoteControlCar();
    }

    public string DistanceDisplay()
    {
        return string.Format($"Driven {_distanceDriven} meters");
    }

    public string BatteryDisplay()
    {
        if (_batteryCharge > 0)
        {
            return string.Format($"Battery at {_batteryCharge}%");
        }
        else
        {
            return string.Format($"Battery empty");
        }
    }

    public void Drive()
    {
        if (_batteryCharge > 0)
        {
            _distanceDriven += 20;
            _batteryCharge -= 1;
        }
    }
}
