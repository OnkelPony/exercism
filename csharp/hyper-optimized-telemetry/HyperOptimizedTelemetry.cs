using System;

public static class TelemetryBuffer
{
	public static byte[] ToBuffer(long reading)
	{
		byte[] result = new byte[9];
		if (reading < int.MinValue || reading > uint.MaxValue)
		{
			result[0] = 256 - 8;
			BitConverter.GetBytes(reading).CopyTo(result, 1);
		}
		else if (reading > int.MaxValue)
		{
			result[0] = 4;
			Array.Copy(BitConverter.GetBytes(reading), 0, result, 1, 4);
		}
		else if (reading < short.MinValue || reading > ushort.MaxValue)
		{
			result[0] = 256 - 4;
			Array.Copy(BitConverter.GetBytes(reading), 0, result, 1, 4);
		}
		else if (reading >= 0)
		{
			result[0] = 2;
			Array.Copy(BitConverter.GetBytes(reading), 0, result, 1, 2);
		}
		else
		{
			result[0] = 256 - 2;
			Array.Copy(BitConverter.GetBytes(reading), 0, result, 1, 2);
		}
		return result;
	}

	public static long FromBuffer(byte[] buffer)
	{
		throw new NotImplementedException("Please implement the static TelemetryBuffer.FromBuffer() method");
	}
}
