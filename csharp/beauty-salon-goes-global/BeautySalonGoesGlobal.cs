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
        return TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), GetTimeZone(location));

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
        const int daysBack = 7;
        TimeZoneInfo timeZone = GetTimeZone(location);
        bool isDstNow = timeZone.IsDaylightSavingTime(dt);
        for (int i = 1; i <= daysBack; i++)
        {
            if (isDstNow != timeZone.IsDaylightSavingTime(dt.AddDays(-i)))
            {
                return true;
            }
        }
        return false;
    }

    public static DateTime NormalizeDateTime(string dtStr, Location location)
    {
		TimeZoneInfo timeZone = GetTimeZone(location);
		string format = timeZone.GetDateTimePattern();
        return DateTime.ParseExact(dtStr, GetTimeZone(location).GetDateTimePattern);
    }

    private static TimeZoneInfo GetTimeZone(Location location)
    {
        TimeZoneInfo timeZone = TimeZoneInfo.Utc;
        if (OperatingSystem.IsWindows())
        {
            switch (location)
            {
                case Location.London:
                    timeZone = TimeZoneInfo.FindSystemTimeZoneById("GMT Standard Time");
                    break;
                case Location.NewYork:
                    timeZone = TimeZoneInfo.FindSystemTimeZoneById("Eastern Standard Time");
                    break;
                case Location.Paris:
                    timeZone = TimeZoneInfo.FindSystemTimeZoneById("W. Europe Standard Time");
                    break;
                default:
                    break;
            }
        }
        else
        {
            switch (location)
            {
                case Location.London:
                    timeZone = TimeZoneInfo.FindSystemTimeZoneById("Europe/London");
                    break;
                case Location.NewYork:
                    timeZone = TimeZoneInfo.FindSystemTimeZoneById("America/New_York");
                    break;
                case Location.Paris:
                    timeZone = TimeZoneInfo.FindSystemTimeZoneById("Europe/Paris");
                    break;
                default:
                    break;
            }
        }
        return timeZone;
    }
}
