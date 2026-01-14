function outer1(x: number): number {
  function inner1(y: number): number {
    const str1: string = "hello";
    console.log(str1);
    return x + y;
  }
  // console.log(str1);
  return inner1(3);
}

console.log(outer1(1));

const outer2 = (x: number): number => {
  const inner2 = (y: number): number => x + y;
  return inner2(x);
};

console.log(outer2(3));

const outer3 = () => {
  let str: string = "hello";
  const inner3 = () => {
    console.log(str);
    str = "hello world";
    console.log(str);
  };
  inner3();
  console.log(str);
};

outer3();

const outer4 = () => {
  const data = { message: "aaa" };
  const inner = () => {
    data.message = "bbb";
  };
  inner();
  console.log(data);
};
outer4();

const outer5 = () => {
  const data = { message: "aaa" } as const
  const inner = () => {
    // data.message = "bbb"; => コンパイルエラー
  };
  inner();
  console.log(data);
};
outer5();
