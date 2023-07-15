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
	public override bool Equals(object obj)
	{
        FacialFeatures other = (FacialFeatures)obj;
		return this.EyeColor == other.EyeColor && this.PhiltrumWidth == other.PhiltrumWidth;
	}

	public override int GetHashCode()
	{
		return EyeColor.GetHashCode() + PhiltrumWidth.GetHashCode();
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
    public override bool Equals(object obj)
    {
        Identity other = (Identity)obj;
        return this.Email == other.Email && this.FacialFeatures.Equals(other.FacialFeatures);
    }

	public override int GetHashCode()
	{
		return Email.GetHashCode() + FacialFeatures.GetHashCode();
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
        if (IsRegistered(identity))
        {
            return false;
        }
        Identities.Add(identity);
        return true;
    }

    public bool IsRegistered(Identity identity)
    {
        return Identities.Contains(identity);
    }

    public static bool AreSameObject(Identity identityA, Identity identityB)
    {
        return identityA == identityB;
    }
}
