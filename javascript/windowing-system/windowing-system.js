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
	constructor() {
		this.screenSize = new Size(800, 600);
		this.size = new Size();
		this.position = new Position();
	}

	resize(newSize) {
		let maxWidth = this.screenSize.width - this.position.x;
		let maxHeight = this.screenSize.height - this.position.y;
		this.size = newSize;
		if (newSize.width > maxWidth) {
			this.size.width = maxWidth;
		}
		if (newSize.height > maxHeight) {
			this.size.height = maxHeight;
		}
		if (newSize.width < 1) {
			this.size.width = 1;
		}
		if (newSize.height < 1) {
			this.size.height = 1;
		}
	}

	move(newPosition) {
		let maxX = this.screenSize.width - this.size.width;
		let maxY = this.screenSize.height - this.size.height;
		this.position = newPosition;
		if (newPosition.x > maxX) {
			this.position.x = maxX;
		}
		if (newPosition.y > maxY) {
			this.position.y = maxY;
		}
		if (newPosition.x < 0) {
			this.position.x = 0;
		}
		if (newPosition.y < 0) {
			this.position.y = 0;
		}
	}
}

/**
 * Implement a changeWindow function that accepts a ProgramWindow instance
 * as input and changes the window to the specified size and position.
 * The function should return the ProgramWindow instance
 * that was passed in after the changes were applied.
 * The window should get a width of 400, a height of 300
 * and be positioned at x = 100, y = 150.
 *
 * @param {ProgramWindow} programWindow
 * @returns {ProgramWindow} passed instance after the changes were applied
 */
export function changeWindow(programWindow) {
	const newPosition = new Position(100, 150);
	const newSize = new Size(400, 300);
	programWindow.move(newPosition);
	programWindow.resize(newSize);
	return programWindow;
}