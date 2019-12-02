
export const intCode = (codes: number[]) => {
    let i = 0;
    for (; i < codes.length; i += 4) {
        let [first, second, third] = [codes[codes[i + 1]], codes[codes[i + 2]], codes[i + 3]]
        switch (codes[i]) {
            case 1:
                codes[third] = first + second
                break;
            case 2:
                codes[third] = first * second
                break;
            case 99:
                break;
        }
    }
    return codes
}

