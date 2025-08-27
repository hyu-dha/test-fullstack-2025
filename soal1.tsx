function factorial(n: number): number {
   if (n === 0 || n === 1) {
      return 1;
   }
   return n * factorial(n - 1);
}

const number: number = 0;
const result: number = factorial(number);
console.log(`Factorial dari ${number} adalah ${result}`);