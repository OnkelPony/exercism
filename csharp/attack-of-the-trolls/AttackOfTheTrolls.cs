using System;

enum AccountType : short
{
    Guest = 1,
    User = 3,
    Moderator = 7
}

[Flags]
enum Permission : short
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
        if (Enum.IsDefined(typeof(AccountType), accountType))
        {
            return (Permission)accountType;
        }
        else
        {
            return Permission.None;
        }
    }

    public static Permission Grant(Permission current, Permission grant)
    {
        return (Permission)((short)current | (short)grant);
    }

    public static Permission Revoke(Permission current, Permission revoke)
    {
        return (Permission)((short)current - revoke);
    }

    public static bool Check(Permission current, Permission check)
    {
        throw new NotImplementedException("Please implement the (static) Permissions.Check() method");
    }
}
