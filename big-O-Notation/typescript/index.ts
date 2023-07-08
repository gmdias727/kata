interface Clock {
    startTime: number;
    endTime: number;
}

interface BigONotation {
    list: number[];
    listSize: number;
    listLength: number; 
}

const clock: Clock = {
    startTime: 0,
    endTime: 0
}

const bigONotation: BigONotation = {
    list: [],
    listSize: 0,
    listLength: 3_000_000,
}

const generateRandomArray = (obj: BigONotation): number[] => {
    for (let index = 0; index < obj.listLength; index++) {
        bigONotation.list[index] = Math.floor(Math.random() * 100000);
    }
    return bigONotation.list;
}

// O(1) - Constant Time
const addItemToArray = (num: number) => {
    return bigONotation.list.push(num);
}

// O(N)
const linearSearch = (num: number) => {
    let valueInArray: boolean = false;
    const indexWithValues: string = '';

    clock.startTime = Date.now();

    for (let index = 0; index < bigONotation.list.length; index++) {
        if (bigONotation.list[index] === num) {
            valueInArray = true;
            indexWithValues.concat(index.toString());
            clock.endTime = Date.now();
            return index;
        }
    }
    
    console.log(`Time taken to linearSearch: ${clock.endTime - clock.startTime}ms`);
    
    if (!valueInArray) {
        return -1;
    }
}

generateRandomArray(bigONotation);
addItemToArray(100);

// console.log("");

linearSearch(300);

