using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace RGBToHex
{
    public class RGBToHEX
    {

        public string Rgb(int r, int g, int b)
        {
            return $"{FixNumber(r):X2}{FixNumber(g):X2}{FixNumber(b):X2}";
        }

        int FixNumber(int number)
        {
            if (number > 255)
            {
                return 255;
            }
            else if ( number <0 )
            {
                return 0;
            }

            return number;
        }



    }
}
