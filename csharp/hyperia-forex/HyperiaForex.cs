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

    // TODO: implement equality operators
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

    // TODO: implement comparison operators

    // TODO: implement arithmetic operators

    // TODO: implement type conversion operators
}
