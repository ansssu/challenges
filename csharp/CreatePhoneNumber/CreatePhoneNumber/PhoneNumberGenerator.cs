using System.Text.RegularExpressions;

namespace CreatePhoneNumber
{
    public class PhoneNumberGenerator
    {

        public string GeneratePhoneNumber(int[] numbers)
        {
            return Regex.Replace(string.Concat(numbers), @"(\d{3})(\d{3})(\d{4})", "($1) $2-$3");
        }

    }
}