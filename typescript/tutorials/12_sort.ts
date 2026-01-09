const numbers: Array<number> = [2, 3, 10, 1];
// sort of dictionary
numbers.sort();
console.log(numbers);

numbers.sort((a, b) => a - b);
console.log(numbers);
numbers.sort((a, b) => b - a);
console.log(numbers);

numbers.push(15);
numbers.sort(() => 0);
console.log(numbers);

// object sort
interface User {
  id: number;
  name: string;
  age: number;
}

const users: User[] = [
  { id: 1, name: "bob", age: 19 },
  { id: 2, name: "boba", age: 16 },
  { id: 3, name: "bobd", age: 18 },
];
users.sort((a, b) => a.age - b.age);
console.log(users);

// japanese sort
const fruits = ["りんご", "ばなな", "田中"];
fruits.sort((a, b) => a.localeCompare(b, "ja"));
console.log(fruits);

// imutable sort
const org = [10, 3, 1, 6];
const sortedOrg = [...org].sort((a, b) => a - b);
console.log(org);
console.log(sortedOrg);
