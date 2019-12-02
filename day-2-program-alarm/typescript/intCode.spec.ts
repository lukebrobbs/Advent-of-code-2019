import { expect } from 'chai';
import { intCode } from './intCode';
import { codes } from './input';

describe("intCode", () => {
    it("IntCode 1 adds numbers together", () => {
        expect(intCode([1, 1, 1, 3])).to.eql([1, 1, 1, 2])
        expect(intCode([1, 0, 0, 0])).to.eql([2, 0, 0, 0])
    })
    it("IntCode 2 multiplies numbers together", () => {
        expect(intCode([2, 1, 1, 3])).to.eql([2, 1, 1, 1])
        expect(intCode([2, 3, 0, 3, 99])).to.eql([2, 3, 0, 6, 99])
    })
    it("IntCode 99 should stop the program", () => {
        expect(intCode([2, 4, 4, 5, 99, 0])).to.eql([2, 4, 4, 5, 99, 9801])
        expect(intCode([1, 1, 1, 4, 99, 5, 6, 0, 99])).to.eql([30, 1, 1, 4, 2, 5, 6, 0, 99])
    })
    it("Should give correct answer to puzzle", () => {
        expect(intCode(codes)[0]).to.equal(5866663)
    })
})

