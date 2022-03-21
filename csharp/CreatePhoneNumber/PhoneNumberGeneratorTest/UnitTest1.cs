using CreatePhoneNumber;
using NUnit.Framework;

namespace PhoneNumberGeneratorTest
{
    public class Tests
    {
        PhoneNumberGenerator? phoneGenerator;

        [SetUp]
        public void Setup()
        {
            phoneGenerator = new PhoneNumberGenerator();
        }

        [Test]
        public void TestPhone()
        {
            Assert.AreEqual("(123) 456-7890", phoneGenerator?.GeneratePhoneNumber(new[] { 1, 2, 3, 4, 5, 6, 7, 8, 9, 0 }));
        }


        public void TestPhoneWithZeros()
        {
            Assert.AreEqual("(000) 000-0000", phoneGenerator?.GeneratePhoneNumber(new[] { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }));
        }
    }
}