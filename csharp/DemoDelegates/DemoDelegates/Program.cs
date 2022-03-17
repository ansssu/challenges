// See https://aka.ms/new-console-template for more information
using DemoLibrary;
using System.Globalization;

var cart = PopulateCarthWithDemoData();


Console.WriteLine($"The total for the cart is {cart.GenerateTotal(SubTotalAlert, CalculateLevelDiscount, AlertUser):C2}");

decimal total = cart.GenerateTotal((subtotal) => Console.WriteLine($"The subtotal for the cart 2 is {subtotal:C2}"),

    (products, subTotal) =>
    {
        if (products.Count > 3)
            return subTotal * 0.5M;
        else
            return subTotal;
    },

    (message) => Console.WriteLine($"Cart 2 Alert:{message}"));
    
    


Console.WriteLine();
Console.WriteLine("Please press any key to exit the application");
Console.ReadKey();
Environment.Exit(0);


static void SubTotalAlert(decimal subTotal)
{
    Console.WriteLine($"The subtotal is {subTotal:C2}");
}

static void AlertUser(string message)
{
    Console.WriteLine(message);
}


static ShoppingCartModel PopulateCarthWithDemoData()
{
    var cart = new ShoppingCartModel();
    cart.Items.Add(new ProductModel { ItemName = "Cereal", Price = 3.63M });
    cart.Items.Add(new ProductModel { ItemName = "Milk", Price = 2.95M });
    cart.Items.Add(new ProductModel { ItemName = "Strawberries", Price = 7.51M });
    cart.Items.Add(new ProductModel { ItemName = "Blueberries", Price = 8.84M });
    return cart;
}


static decimal CalculateLevelDiscount(List<ProductModel> items, decimal subTotal)
{
    if (subTotal > 100)
    {
        return subTotal * 0.80M;
    }
    else if (subTotal > 50)
    {
        return subTotal * 0.85M;
    }
    else if (subTotal > 10)
    {
        return subTotal * 0.90M;
    }
    else
    {
        return subTotal;
    }
}
