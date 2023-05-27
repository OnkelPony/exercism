using System;

class RemoteControlCar
{
    public int _speed;
    public int _batteryDrain;
    public int _distanceDriven;
    public int _batteryCharge;

    // TODO: define the constructor for the 'RemoteControlCar' class
    public RemoteControlCar(int speed, int batteryDrain)
    {
        _batteryDrain = batteryDrain;
        _speed = speed;
        _batteryCharge = 100;
        _distanceDriven = 0;
    }
    public bool BatteryDrained()
    {
        return _batteryCharge < _batteryDrain;
    }

    public int DistanceDriven()
    {
        return _distanceDriven;
    }

    public void Drive()
    {
        if (_batteryCharge >= _batteryDrain) {
        _batteryCharge -= _batteryDrain;
        _distanceDriven += _speed;
        }
    }

    public static RemoteControlCar Nitro()
    {
        return new RemoteControlCar(50, 4);
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
        return (float)_distance / car._speed <= car._batteryCharge / car._batteryDrain;
    }
}
