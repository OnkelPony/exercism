using System;

public class Orm
{
    private Database database;

    public Orm(Database database)
    {
        this.database = database;
    }

    public void Begin()
    {
        if (database.DbState != Database.State.Closed)
        {
            throw new InvalidOperationException();
        }
        database.BeginTransaction();
    }

    public void Write(string data)
    {
        if (database.DbState != Database.State.TransactionStarted || data == "bad write")
        {
            database.EndTransaction();
            throw new InvalidOperationException();
        }
        database.Write(data);
    }

    public void Commit()
    {
        if (database.DbState!= Database.State.DataWritten)
        {
            database.EndTransaction();
            throw new InvalidOperationException();
        }
        database.Dispose();
    }
}
