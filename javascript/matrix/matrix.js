//
// This is only a SKELETON file for the 'Matrix' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class Matrix {
    constructor(numberString) {
        this.result = [];
        this.matrixRows = numberString.split("\n");
        this.matrixRows.forEach(row => this.result.push(row.split(" ").map(element => Number(element))));
    }

    get rows() {
        return this.result;
    }

    get columns() {
        return this.result[0].map((_, i) => this.result.map(row => row[i]));
    }
}
