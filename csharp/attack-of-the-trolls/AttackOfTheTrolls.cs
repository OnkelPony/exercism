using System;

enum AccountType : byte
{
    Guest = 1,
    User = 3,
    Moderator = 7
}

[Flags]
enum Permission : byte
{
    None = 0,
    Read = 1,
    Write = 2,
    Delete = 4,
    All = 7,
}
static class Permissions
{
    public static Permission Default(AccountType accountType)
    {
        return Enum.IsDefined(typeof(AccountType), accountType) ? (Permission)accountType : Permission.None;

    }

    public static Permission Grant(Permission current, Permission grant)
    {
        return (Permission)((byte)current | (byte)grant);
    }

    public static Permission Revoke(Permission current, Permission revoke)
    {
        return (Permission)((byte)current - revoke);
    }

    public static bool Check(Permission current, Permission check)
    {
        return ((short)current | (short)check) == (short)current;
    }
}
