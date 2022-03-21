using System.Numerics;
using System.Text;

namespace NumberOfTraillingZerosOfN
{
    public class TrailingZerosCounter
    {

        
        public int CountZeros(int number)
        {
            if (number < 0) return -1;
            int fives =0;
            for (int i = 5; number / i >= 1; i *= 5)
                fives += number / i;
            return fives;
        }
    }
}