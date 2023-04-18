export class Matrix {
    constructor(numberString) {
        this.result = numberString.split("\n")
            .map(line => line.split(" ")
                .map(Number));
    }

    static transMatrix(rows) {
        return rows[0].map((_, i) => rows.map(row => row[i]));
    }
    get rows() {
        return this.result;
    }

    get columns() {
        return Matrix.transMatrix(this.result);
    }
}
