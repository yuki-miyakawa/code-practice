const v1: string = "hello";
const v2: string = "Hello";
const v3: string = "hello";

if (v1 === v2) {
  console.log("v1 is the same v2: {v1, v2}");
} else if (v1 === v3) {
  console.log("v1 is the same v3, v1: %s, v3: %s", v1, v3);
} else {
  console.log("different");
}

const n7 = 1;
const n8 = 2;

if (n7 > 0 && n8 > 0) {
  console.log(n7, n8);
}


