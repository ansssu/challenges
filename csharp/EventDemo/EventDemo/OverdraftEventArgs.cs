using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace EventDemo
{
    public class OverdraftEventArgs: EventArgs
    {
        public decimal AmountOverdrafted { get; private set;} 
        public string? MoreInfo { get; private set; }
        public bool CancelTransaction { get; set; }

        public OverdraftEventArgs(decimal amountOverdrafted, string moreInfo)
        {
            AmountOverdrafted = amountOverdrafted;
            MoreInfo = moreInfo;
        }
    }
}
