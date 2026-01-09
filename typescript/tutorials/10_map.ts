const map2 = new Map<string, number>();
console.log(map2);
map2.set("apple", 3);
console.log(map2);
console.log(map2.get("apple"));
console.log(map2.get("banana"));
console.log(map2.has("apple"));
console.log(map2.has("banana"));

map2.set("delete", 5);
console.log(map2);
map2.delete("delete");
console.log(map2);

const map3 = new Map<number, string>();
console.log(map3.set(1, "apple"));
console.log(map3.delete(1));
console.log(map3);

const map4 = new Map<string, string>();

map4.set("apple", "りんご");
map4.set("banana", "ばなな");
map4.set("orange", "みかん");
for (const [key, value] of map4) {
  console.log(key, value);
}

for (const key of map4.keys()) {
  console.log(key);
}
for (const value of map4.values()) {
  console.log(value);
}

console.log(map4.size);
