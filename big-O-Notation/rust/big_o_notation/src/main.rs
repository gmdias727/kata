struct BigONotation {
    the_array: Vec<i32>,
    array_size: i32,
    items_in_array: i32,
}

impl BigONotation {
    fn new() -> Self {
        Self {
            the_array: Vec::new(),
            array_size: 0,
            items_in_array: 0,
        }
    }

    fn add_item_to_array(&mut self, new_item: i32) {
        self.the_array.push(new_item);
        self.array_size = self.the_array.len() as i32;
        self.items_in_array += 1;
        // println!("{}", new_item);
        println!("{:?}", self.the_array);
    }
}


fn main()   {

    let mut my_notation : BigONotation = BigONotation::new();
    my_notation.add_item_to_array(3);
    println!("{:?}", my_notation.add_item_to_array(6));
    println!("{:?}", my_notation.add_item_to_array(9));
}
