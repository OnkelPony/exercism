using System;
public struct CurrencyAmount
{
    private decimal amount;
    private string currency;

    public CurrencyAmount(decimal amount, string currency)
    {
        this.amount = amount;
        this.currency = currency;
    }

    public static bool operator ==(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException("Currency mismatch");
        }
        return left.amount == right.amount;
    }

    public static bool operator !=(CurrencyAmount left, CurrencyAmount right)
    {
        return !(left == right);
    }

    public static bool operator >(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException("Currency mismatch");
        }
        return left.amount > right.amount;
    }

    public static bool operator <(CurrencyAmount left, CurrencyAmount right)
    {
        return !(left > right);
    }
    // TODO: implement arithmetic operators
    public static CurrencyAmount operator +(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException("Currency mismatch");
        }
        return new CurrencyAmount(left.amount + right.amount, left.currency);
    }

    public static CurrencyAmount operator -(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException("Currency mismatch");
        }
        return new CurrencyAmount(left.amount - right.amount, left.currency);
    }

    // TODO: implement type conversion operators
}
