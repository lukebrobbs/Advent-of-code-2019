import { modules } from "./input";

const calculateFuelForModule = (module: number) => {
  const num = Math.floor(module / 3 - 2);
  let answer = num >= 0 ? num : 0;
  if (num > 0) {
    answer += calculateFuelForModule(num);
  }
  return answer;
};

const answer = modules.reduce(
  (acc, curr) => acc + calculateFuelForModule(curr),
  0
);

console.log(answer);
