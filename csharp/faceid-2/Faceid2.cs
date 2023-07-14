using System;
using System.Collections.Generic;

public class FacialFeatures
{
    public string EyeColor { get; }
    public decimal PhiltrumWidth { get; }

    public FacialFeatures(string eyeColor, decimal philtrumWidth)
    {
        EyeColor = eyeColor;
        PhiltrumWidth = philtrumWidth;
    }
	// TODO: implement equality and GetHashCode() methods
	public  bool Equals(FacialFeatures other)
	{
		return this.EyeColor == other.EyeColor && this.PhiltrumWidth == other.PhiltrumWidth;
	}
}

public class Identity
{
    public string Email { get; }
    public FacialFeatures FacialFeatures { get; }

    public Identity(string email, FacialFeatures facialFeatures)
    {
        Email = email;
        FacialFeatures = facialFeatures;
    }
    // TODO: implement equality and GetHashCode() methods
    public bool Equals(Identity other)
    {
        return this.Email == other.Email && this.FacialFeatures == other.FacialFeatures;
    }
}

public class Authenticator
{
    private HashSet<Identity> Identities = new HashSet<Identity>();
    public static bool AreSameFace(FacialFeatures faceA, FacialFeatures faceB)
    {
        return faceA.Equals(faceB);
    }

    public bool IsAdmin(Identity identity)
    {
        return identity.Equals(new Identity("admin@exerc.ism", new FacialFeatures("green", 0.9m)));
    }

    public bool Register(Identity identity)
    {
        if (Identities.Contains(identity))
        {
            return false;
        }
        Identities.Add(identity);
        return true;
    }

    public bool IsRegistered(Identity identity)
    {
        throw new NotImplementedException("Please implement the Authenticator.IsRegistered() method");
    }

    public static bool AreSameObject(Identity identityA, Identity identityB)
    {
        throw new NotImplementedException("Please implement the Authenticator.AreSameObject() method");
    }
}
