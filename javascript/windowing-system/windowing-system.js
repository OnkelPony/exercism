// @ts-check

/**
 * Define Size for storing the dimensions of the window.
 *
 * @param {number} width
 * @param {number} height
 */
export function Size(width = 80, height = 60) {
	this.width = width;
	this.height = height;
}

/**
 * Takes a new width and height as parameters
 * and changes the fields to reflect the new size.
 *
 * @param {number} newWidth
 * @param {number} newHeight
 */
Size.prototype.resize = function (newWidth, newHeight) {
	this.width = newWidth;
	this.height = newHeight;
};

/**
 * Define Position to store a window position.
 *
 * @param {number} x;
 * @param {number} y;
 */
export function Position(x = 0, y = 0) {
	this.x = x;
	this.y = y;
}

/**
 * Takes new x and y parameters and changes the properties
 * to reflect the new position.
 *
 * @param {number} newX
 * @param {number} newY
 */
Position.prototype.move = function (newX, newY) {
	this.x = newX;
	this.y = newY;
}

/**
 * Define a ProgramWindow class.
 */
export class ProgramWindow {
	constructor () {
		this.screenSize = new Size(800, 600);
		this.size = new Size();
		this.position = new Position();
	}
}