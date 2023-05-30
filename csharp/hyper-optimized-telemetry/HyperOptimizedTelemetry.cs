using System;

public static class TelemetryBuffer
{
	public static byte[] ToBuffer(long reading)
	{
        byte[] result = new byte[9];
        if (reading <= int.MinValue || reading > uint.MaxValue)
        {
        result[0] = 256 - 8;
        }
        else if (reading <=
		 BitConverter.GetBytes(reading).CopyTo(result, 1);
         return result;
	}

	public static long FromBuffer(byte[] buffer)
	{
		throw new NotImplementedException("Please implement the static TelemetryBuffer.FromBuffer() method");
	}
}
