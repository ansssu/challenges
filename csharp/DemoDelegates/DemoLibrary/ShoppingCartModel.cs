﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace DemoLibrary
{

    public class ShoppingCartModel
    {

        public delegate void MentionDiscount(decimal subTotal); 

        public List<ProductModel> Items { get; set; } = new List<ProductModel>();

        public decimal GenerateTotal(MentionDiscount mentionSubTotal,
                                    Func<List<ProductModel>, decimal, decimal> calculateDiscountTotal
                                    Action<string> telUseWereDiscounting
            )
        {
            decimal subTotal = Items.Sum(x => x.Price);

            mentionSubTotal(subTotal);
            telUseWereDiscounting("We are applying your discount");

            return calculateDiscountTotal(Items, subTotal);

        }

       

    }
}
