using MaxSumContiguousArray;
using NUnit.Framework;

namespace MaxSumContiguousArrayTest
{
    public class Tests
    {

        KadaneAlgo kadaneAlgo;

        [SetUp]
        public void Setup()
        {
            kadaneAlgo = new KadaneAlgo();
        }

        [Test]
        public void BasicTests()
        {
            Assert.AreEqual(6, kadaneAlgo.RunKadane(new[] { -2, 1, -3, 4, -1, 2, 1, -5, 4 }));
        }
    }
}