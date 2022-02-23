// See https://aka.ms/new-console-template for more information



Console.WriteLine("Type a number:");

var number = 0;
var sumOfNumbers = 0;
checked
{
    number = Convert.ToInt32(Console.ReadLine());
}



for (int i = number-1; i >= 0; i--)
{
   if (i % 3 == 0 || i % 5 == 0) {
       sumOfNumbers+=i;
   }
}

Console.WriteLine($"Sum of numbers wich are multiple of 3 or 5: {sumOfNumbers} ") ;




