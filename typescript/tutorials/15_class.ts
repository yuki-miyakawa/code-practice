class User {
  name: string;
  age: number;

  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
  }

  greet(): void {
    console.log(`hello, ${this.name}`);
  }
}

const user1 = new User("john", 5);
user1.greet();
