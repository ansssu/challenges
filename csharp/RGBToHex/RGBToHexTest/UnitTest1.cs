using NUnit.Framework;
using RGBToHex;

namespace RGBToHexTest
{
    [TestFixture]
    public class Tests
    {

        RGBToHEX rgbParse;

        [SetUp]
        public void Setup()
        {
            rgbParse = new RGBToHEX();
        }

        [Test]
        public void FixedTests()
        {
            Assert.AreEqual("FFFFFF", rgbParse.Rgb(255, 255, 255));
            Assert.AreEqual("FFFFFF", rgbParse.Rgb(255, 255, 300));
            Assert.AreEqual("000000", rgbParse.Rgb(0, 0, 0));
            Assert.AreEqual("9400D3", rgbParse.Rgb(148, 0, 211));
            Assert.AreEqual("9400D3", rgbParse.Rgb(148, -20, 211), "Handle negative numbers.");
            Assert.AreEqual("90C3D4", rgbParse.Rgb(144, 195, 212));
            Assert.AreEqual("D4350C", rgbParse.Rgb(212, 53, 12), "Consider single hex digit numbers.");
        }
    }
}