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
            Func<int, string> f = x => (x < 0 ? 0 : x > 255 ? 255 : x).ToString("X2");
            return $"{f(r)}{f(g)}{f(b)}";
        }
    }
}
