using System;
using System.Collections.Generic;
using Microsoft.VisualBasic;

public struct Coord
{
    public Coord(ushort x, ushort y)
    {
        X = x;
        Y = y;
    }

    public ushort X { get; }
    public ushort Y { get; }
}

public struct Plot : IComparable<Plot>
{
    public Plot(Coord A, Coord B, Coord C, Coord D)
    {
        this.A = A;
        this.B = B;
        this.C = C;
        this.D = D;
    }

    public Coord A { get; }
    public Coord B { get; }
    public Coord C { get; }
    public Coord D { get; }

    public int CompareTo(Plot other)
    {
        return (int)(this.LongestSide() - other.LongestSide());
    }

    public double LongestSide()
    {
        double a = Math.Sqrt(Math.Pow(B.X - A.X, 2) + Math.Pow(B.Y - A.Y, 2));
        double b = Math.Sqrt(Math.Pow(C.X - B.X, 2) + Math.Pow(C.Y - B.Y, 2));
        double c = Math.Sqrt(Math.Pow(D.X - C.X, 2) + Math.Pow(D.Y - C.Y, 2));
        double d = Math.Sqrt(Math.Pow(A.X - D.X, 2) + Math.Pow(A.Y - D.Y, 2));
        return Math.Max(a, Math.Max(b, Math.Max(c, d)));
    }
}


public class ClaimsHandler
{
    List<Plot> claims = new List<Plot>();
    public void StakeClaim(Plot plot)
    {
        claims.Add(plot);
    }

    public bool IsClaimStaked(Plot plot)
    {
        return claims.Contains(plot);
    }

    public bool IsLastClaim(Plot plot)
    {
        return claims[claims.Count - 1].Equals(plot);
    }

    public Plot GetClaimWithLongestSide()
    {
        Plot longestSidePlot = claims[0];
        foreach (Plot claim in claims)
        {
            if (claim.CompareTo(longestSidePlot) > 0)
            {
                longestSidePlot = claim;
            }
        }
        return longestSidePlot;
    }
}
