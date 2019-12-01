const { modules } = require('./input')

const calculateFuelForModule = (n) => {
    const num = Math.floor(n / 3 - 2);
    let answer = num >= 0 ? num : 0
    if (num > 0) {
        answer += f(num)
    }
    return answer;
}

const answer = modules.reduce((acc, curr) => acc + calculateFuelForModule(curr), 0)

console.log(answer)