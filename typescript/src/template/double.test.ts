import {double} from './double';
// bun test src/template/double.test.ts

describe('double', () =>{
  test('double of 3 is 6', () => {
    expect(double(3)).toBe(3);
  });
})
