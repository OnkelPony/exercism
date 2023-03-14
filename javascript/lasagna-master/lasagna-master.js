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
	for (const element of layers) {
		if (element === "sauce") {
			sauceAmount += SAUCE_PER_LAYER;
		} else if (element === "noodles") {
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
 *
 * @param {recipe} recipe
 * @param {number} portions
 * @return {recipe} scaled recipe
 */
export function scaleRecipe(recipe, portions) {
	let scaledRecipe = {};
	for (let ingredient in recipe) {
		scaledRecipe[ingredient] = recipe[ingredient] / RECIPE_PORTIONS * portions;
	}
	return scaledRecipe;
}