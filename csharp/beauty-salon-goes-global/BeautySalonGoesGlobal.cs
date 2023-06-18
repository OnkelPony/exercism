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
        switch (alertLevel)
        {
            case AlertLevel.Early: return appointment.AddDays(-1);
            case AlertLevel.Standard: return appointment.AddHours(-1).AddMinutes(-45);
            case AlertLevel.Late: return appointment.AddMinutes(-30);
            default: return appointment;
        }
    }

    public static bool HasDaylightSavingChanged(DateTime dt, Location location)
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

    public static DateTime NormalizeDateTime(string dtStr, Location location)
    {
        throw new NotImplementedException("Please implement the (static) Appointment.NormalizeDateTime() method");
    }
}
