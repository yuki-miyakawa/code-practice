const arr1: string[] = ["a", "b"];
console.log(arr1);

const arr2: Array<string> = ["c", "d"];
console.log(arr2);

const arr3: (string | number)[] = ["e", 2];
console.log(arr3);

let arr4: [string, number];
arr4 = ["f", 41];
console.log(arr4);

const arr5: readonly number[] = [1, 2, 34];
console.log(arr5);

const arr6: Array<string | number> = [
  1,
  "two",
  3,
  "four",
  5,
  "six",
  7,
  "eight",
];
console.log(arr6);

console.log(arr6.slice(3, 5));

// append
arr6.push(9);
console.log(arr6);

const arr7 = [...arr6, "ten"];
console.log(arr7);

arr6.pop();
console.log(arr6);
console.log(arr7);
