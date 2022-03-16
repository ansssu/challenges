   public static class LikeTranslator
    {

        public static string Translate(string[] names)
        {
            string[] pattern = {
                    "no one likes this",
                    "{0} likes this ",
                    "{0} and {1} like this",
                    "{0}, {1} and {2} like this ",
                    "{0}, {1} and {2} others like this."
             };

            string message = names.Length < 4 ? string.Format(pattern[names.Length], names) : string.Format(pattern[4], names[0], names[1], names.Length - 2);
            return message;
        }

    }
