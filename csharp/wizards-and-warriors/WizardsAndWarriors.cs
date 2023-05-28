using System;

abstract class Character
{
    protected string _characterType;
    protected bool _isVulnerable = false;
    protected int _damagePoints;
    protected Character(string characterType)
    {
        _characterType = characterType;
    }

    public abstract int DamagePoints(Character target);

    public virtual bool Vulnerable()
    {
        return _isVulnerable;
    }

    public override string ToString()
    {
        return $"Character is a {_characterType}";
    }
}

class Warrior : Character
{
    public Warrior() : base("Warrior")
    {
    }

    public override int DamagePoints(Character target)
    {
        if (target.Vulnerable())
        {
            return 10;
        }
        else
        {
            return 6;
        }
    }
}

class Wizard : Character
{
    protected bool _spellReady = false;
    public Wizard() : base("Wizard")
    {
        base._isVulnerable = true;
        _damagePoints = 3;
    }

    public override int DamagePoints(Character target)
    {
        return _damagePoints;
    }

    public void PrepareSpell()
    {
        _spellReady = true;
        _isVulnerable = false;
        _damagePoints = 12;
    }
}
