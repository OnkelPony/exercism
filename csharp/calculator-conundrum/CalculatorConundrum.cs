using System;

public static class SimpleCalculator
{
    public static string Calculate(int operand1, int operand2, string operation)
    {
        if (operand2 == 0)
        {
            return "Division by zero is not allowed.";
        }
        switch (operation)
        {
            case "":
                throw new ArgumentException();
            case null:
                throw new ArgumentNullException();
            case "+":
                return $"{operand1} {operation} {operand2} = {SimpleOperation.Addition(operand1, operand2)}";
            case "*":
                return $"{operand1} {operation} {operand2} = {SimpleOperation.Multiplication(operand1, operand2)}";
            case "/":
                return $"{operand1} {operation} {operand2} = {SimpleOperation.Division(operand1, operand2)}";
            default:
                throw new ArgumentOutOfRangeException();
        }
    }
}
