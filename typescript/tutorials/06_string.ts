const str1: string = "hello";
console.log(str1.length);

const str2: string = "こんにちは";
console.log(str2.length);

const str3: string = "satoh";
console.log(`ようこそ${str3}さま`);
console.log("ようこそ%sさま", str3);

const str4: string = "hello world";
console.log(str4.slice(2));
console.log(str4.slice(2, 5));
console.log(str4.slice(-5));

// strconv
const str5: string = "1234";
const strnum: number = Number(str5);
console.log(strnum);
const str6: string = "1234a";
const strnum2: number = Number(str6);
console.log(strnum2);

// join
const strArr: Array<string> = ["aaa", "bbb", "ccc"];
console.log(strArr.join(","));

// split
const str7: string = "hello, morning, evening";
const strArr2: string[] = str7.split(",");
console.log(strArr2);
