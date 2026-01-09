function add1(a: number, b: number): number {
  return a + b;
}
console.log(add1(3, 5));

const add2 = (a: number, b: number): number => {
  return a + b;
};
console.log(add2(5, 8));

const add3 = (a: number, b: number): number => a + b;
console.log(add3(4, 4));

const strFun1 = (name: string, title?: string): string => {
  if (title) {
    return `Hello, ${title}, ${name}`;
  }
  return `Hello, ${name}`;
};
console.log(strFun1("taro"));
console.log(strFun1("jiro", "world"));

const strFun2 = (name: string, title: string = "www"): string => {
  return `Hello, ${title}, ${name}`;
};
console.log(strFun2("saburo"));
console.log(strFun2("shiro", "qqq"));

const testFun1 = (...numbers: number[]): number[] => {
  return numbers;
};
console.log(testFun1(1, 2, 2));

// 関数の型定義以降(https://gemini.google.com/app/609f96c0e29f99d7)
