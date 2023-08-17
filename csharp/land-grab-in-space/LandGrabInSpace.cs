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

public struct Plot
{
    public Plot(Coord tl, Coord tr, Coord bl, Coord br)
    {
        TL = tl;
        TR = tr;
        BL = bl;
        BR = br;
    }

    public Coord TL { get; }
    public Coord TR { get; }
    public Coord BL { get; }
    public Coord BR { get; }
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
        throw new NotImplementedException("Please implement the ClaimsHandler.IsLastClaim() method");
    }

    public Plot GetClaimWithLongestSide()
    {
        throw new NotImplementedException("Please implement the ClaimsHandler.GetClaimWithLongestSide() method");
    }
}
