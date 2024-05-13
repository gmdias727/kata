const x = 1 + 2;

const add = (a: number, b: number) => a + b;
const addx = (a: number) => (b: number) => a + b;

const three = add(1, 2);
const threex = addx(1)(2);

const messageSize = (n: number) => {
  const elementSize = 4;
  const separatorSize = 1;
  const constant = elementSize + separatorSize;
  return constant * n;
};

const printConstant = (constant: number) => {
  console.log(`constant: ${constant}`);
}

const _ = messageSize(
  (() => {
    const base = 5;
    return base + 6;
  })()
);

type Status = | "valid" | "banned";
const user_status: Status = "valid";