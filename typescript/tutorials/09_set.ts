const mySet = new Set<string>();

mySet.add("A");
mySet.add("B");
mySet.add("B");
mySet.add("C");

console.log(mySet);

console.log(mySet.has("A"));
console.log(mySet.has("D"));
console.log(mySet.size);

const isDelete = mySet.delete("A");
console.log(isDelete);
console.log(mySet);

const arr9: Array<number> = [1, 2, 2, 3, 4, 4, 5, 5, 6, 6];
console.log(arr9);

const arr10 = [...new Set(arr9)];
console.log(arr10);
