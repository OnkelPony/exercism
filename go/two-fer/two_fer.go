/*
Package twofer implements method, which returns string, depending on parameter.

*/
package twofer

// Returns One for X, one for me., where X is the parameter. If parameter is missing, it pastes "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
