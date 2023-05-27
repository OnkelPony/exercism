using System;

class RemoteControlCar
{
    private int _speed;
    private int _batteryDrain;
    private int _distanceDriven;
    private int _batteryCharge;

    // TODO: define the constructor for the 'RemoteControlCar' class
    public RemoteControlCar(int speed, int batteryDrain)
    {
        _batteryDrain = batteryDrain;
        _speed = speed;
    }
    public bool BatteryDrained()
    {
        throw new NotImplementedException("Please implement the RemoteControlCar.BatteryDrained() method");
    }

    public int DistanceDriven()
    {
        throw new NotImplementedException("Please implement the RemoteControlCar.DistanceDriven() method");
    }

    public void Drive()
    {
        _batteryCharge -= _batteryDrain;
        _distanceDriven += _speed;
    }

    public static RemoteControlCar Nitro()
    {
        throw new NotImplementedException("Please implement the (static) RemoteControlCar.Nitro() method");
    }
}

class RaceTrack
{
    private int _distance;

    // TODO: define the constructor for the 'RaceTrack' class
    public RaceTrack(int distance)
    {
        _distance = distance;
    }
    public bool TryFinishTrack(RemoteControlCar car)
    {
        throw new NotImplementedException("Please implement the RaceTrack.TryFinishTrack() method");
    }
}
