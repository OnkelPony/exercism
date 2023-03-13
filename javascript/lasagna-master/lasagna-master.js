/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Grams of noodles needed for lasagna layer.
 */
export const NOODLES_PER_LAYER = 50;

/**
 * Liters of sauces needed for lasa layer.
 */
export const SAUCE_PER_LAYER = 0.2;

/**
 * Portions in recipe
 */
export const RECIPE_PORTIONS = 2;

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
 *
 * @param {string[]} layers
 * @param {number} minutesPerLayer
 * @return {number} preparation time
 */
export function preparationTime(layers, minutesPerLayer = 2) {
	return minutesPerLayer * (layers.length);
}

/**
 * Compute the amounts of noodles and sauce needed.
 *
 * @typedef {Object} amounts
 * @property {number} noodles
 * @property {number} sauce
 *
 * @param {string[]} layers
 * @return {amounts}
 */
export function quantities(layers) {
	let noodlesAmount = 0;
	let sauceAmount = 0;
	for (let i = 0; i < layers.length; i++) {
		if (layers[i] === "sauce") {
			sauceAmount += SAUCE_PER_LAYER;
		} else if (layers[i] === "noodles") {
			noodlesAmount += NOODLES_PER_LAYER;
		}
	}
	return { noodles: noodlesAmount, sauce: sauceAmount };
}

/**
 * Add the secret ingredient.
 *
 * @param {String[]} friensList
 * @param {String[]} myList
 */
export function addSecretIngredient(friensList, myList) {
	let secretIngredient = friensList[friensList.length - 1];
	myList.push(secretIngredient);
}

/**
 * Scale the recipe.
 *
 * @typedef {Object} recipe
 * @property {number} noodles
 * @property {number} sauce
 * @property {number} mozarella
 * @property {number} meat
 *
 * @param {recipe} recipe
 * @param {number} portions
 * @return {recipe} scaled recipe
 */
export function scaleRecipe(recipe, portions) {
	let scaledNoodles = recipe.noodles / RECIPE_PORTIONS * portions;
	let scaledSauce = recipe.sauce / RECIPE_PORTIONS * portions;
	let scaledMozarella = recipe.mozarella / RECIPE_PORTIONS * portions;
	let scaledMeat = recipe.meat / RECIPE_PORTIONS * portions;
	return {
		noodles: scaledNoodles,
		sauce: scaledSauce,
		mozarella: scaledMozarella,
		meat: scaledMeat,
	}
}