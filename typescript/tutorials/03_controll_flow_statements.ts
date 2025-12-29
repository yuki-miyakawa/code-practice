for (let i: number = 0; i < 10; i++) {
  console.log(i);
}

let j: number = 0;
while (j < 10) {
  console.log(j);
  j++;
  if (j == 5) {
    continue;
  }
}
