class Program
{
    public static List<int> the_array = new();
    public static int array_size;
    public static int items_in_array;

    public static void add_item_to_array(int new_item)
    {
        the_array.Add(new_item);
    }

    public static void Main()
    {
        Console.WriteLine(add_item_to_array(10));
    }
}


