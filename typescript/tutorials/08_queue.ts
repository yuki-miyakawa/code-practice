const stack: string[] = [];
stack.push("A");
stack.push("B");
stack.push("C");

console.log(stack);

const last1 = stack.pop();
console.log(last1);
console.log(stack);

const queue: string[] = [];
queue.push("A");
queue.push("B");
queue.push("C");

console.log(queue);

const firstItem = queue.shift();
console.log(firstItem);
console.log(queue);
