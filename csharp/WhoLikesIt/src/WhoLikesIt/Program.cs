// See https://aka.ms/new-console-template for more information



string[] names = new string[3];
string likes = Translate(names);
Console.WriteLine(likes);


string Translate(string[] likes){

     List<(int,string)> translatorTemplates   = new List<(int , string )>{
        (0,"no one likes this"),
        (1,"{0} likes this "),
        (2,"{0} and {1} like this"),
        (3,"{0}, {1} and {2} like this "),
        (4,"{0}, {1} and {2} others like this.")
     };

    int totalLikes =  likes == null ? 0  : likes.Where(x=>!string.IsNullOrEmpty(x)).Count();

    int templateSelector = totalLikes > translatorTemplates.Count-1 ? translatorTemplates.Count-1 : totalLikes;
    string template = translatorTemplates.Where(x=>x.Item1 >= templateSelector ).FirstOrDefault().Item2;
    int arrayCapacity = totalLikes > 3 ? 3 : totalLikes;
    object[] names = likes.AsEnumerable().Select((name,index)=> new {name,index}).Where(x=>x.index < 3).Select(x=>x.name).ToArray();
    if (totalLikes >3)
        names[arrayCapacity-1] =  (totalLikes -2).ToString();
    
    return string.Format(template, names);
        

}

