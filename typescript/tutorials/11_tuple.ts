let user: [string, number];
user = ["alice", 24];
console.log(user);
console.log(user[0].toUpperCase());
console.log(user[1].toFixed(1));
console.log(user[2]);

// labeled tuple
let user2: [name: string, age: number];
user2 = ["bob", 30];
console.log(user2);

let user3: [number, number, number?];
user3 = [1, 2, 3];
console.log(user3);
user3 = [1, 2];
console.log(user3);

// read only
const user4: readonly [number, number] = [10, 30];
user4[0] = 20;
console.log(user4);
const user5: [number, number] = [10, 30];
user5[0] = 20;
console.log(user5);
