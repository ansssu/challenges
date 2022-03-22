namespace MaxSumContiguousArray
{
    public class KadaneAlgo
    {

        public int RunKadane(int [] arr)
        {

            int max_so_far = 0 , max_ending_here = 0;

            foreach (var arrayItem in arr)
            {
                max_ending_here += arrayItem;
                if (max_ending_here < 0) max_ending_here = 0;
                max_so_far = Math.Max(max_ending_here, max_so_far);
            }

            return max_so_far;
        }

    }
}