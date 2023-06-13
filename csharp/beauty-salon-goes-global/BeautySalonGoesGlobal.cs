using System;


public enum Location
{
    NewYork,
    London,
    Paris
}

public enum AlertLevel
{
    Early,
    Standard,
    Late
}

public static class Appointment
{
    public static DateTime ShowLocalTime(DateTime dtUtc)
    {
        return TimeZoneInfo.ConvertTimeFromUtc(dtUtc, TimeZoneInfo.Local);
    }

    public static DateTime Schedule(string appointmentDateDescription, Location location)
    {
        if (OperatingSystem.IsWindows())
        {
            switch (location)
            {
                case Location.London:
                    return TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), TimeZoneInfo.FindSystemTimeZoneById("GMT Standard Time"));
                case Location.NewYork:
                    return TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), TimeZoneInfo.FindSystemTimeZoneById("Eastern Standard Time"));
                case Location.Paris:
                    return TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), TimeZoneInfo.FindSystemTimeZoneById("W. Europe Standard Time"));
                default:
                    return DateTime.Parse(appointmentDateDescription);
            }
        }
        else
        {
            switch (location)
            {
                case Location.London:
                    return TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), TimeZoneInfo.FindSystemTimeZoneById("Europe/London"));
                case Location.NewYork:
                    return TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), TimeZoneInfo.FindSystemTimeZoneById("America/New_York"));
                case Location.Paris:
                    return TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), TimeZoneInfo.FindSystemTimeZoneById("Europe/Paris"));
                default:
                    return DateTime.Parse(appointmentDateDescription);
            }
        }
    }
    public static DateTime GetAlertTime(DateTime appointment, AlertLevel alertLevel)
    {
        throw new NotImplementedException("Please implement the (static) Appointment.GetAlertTime() method");
    }

    public static bool HasDaylightSavingChanged(DateTime dt, Location location)
    {
        throw new NotImplementedException("Please implement the (static) Appointment.HasDaylightSavingChanged() method");
    }

    public static DateTime NormalizeDateTime(string dtStr, Location location)
    {
        throw new NotImplementedException("Please implement the (static) Appointment.NormalizeDateTime() method");
    }
}
