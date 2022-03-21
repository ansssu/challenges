using NumberOfTraillingZerosOfN;
using NUnit.Framework;

namespace NumberOfTraillingZerosOfNTest
{
    public class Tests
    {

        TrailingZerosCounter? countTraillingZeros;

        [SetUp]
        public void Setup()
        {
            countTraillingZeros = new TrailingZerosCounter();
        }

        [Test]
        public void BasicTest()
        {
            Assert.AreEqual(1, countTraillingZeros?.CountZeros(6));
            Assert.AreEqual(2, countTraillingZeros?.CountZeros(12));
        }
    }
}