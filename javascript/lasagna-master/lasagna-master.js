/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

/**
 * Determins the status of the lasagna.
 *
 * @param {number} minutesRemaining
 * @return {string} lasagna status
 */
export function cookingStatus(minutesRemaining) {
	switch (minutesRemaining) {
		case 0:
			return "Lasagna is done.";
		case undefined:
			return "You forgot to set the timer.";
		default:
			return "Not done, please wait.";
	}
}

/**
 * Estimate the preparation time.
 * @param {string[]} layers
 * @param {number} minutesPerLayer
 * @return {number} preparation time
 */
export function preparationTime(layers, minutesPerLayer = 2) {
	return minutesPerLayer * (layers.length);
}