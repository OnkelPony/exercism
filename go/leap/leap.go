/*
Package leap implements check for leap years
 */
package leap

const four = 4
const hundred = 100
const fourHundred = 400

// Returns if year provided in parameter is leap or not
func IsLeapYear(year int ) bool {

	if year % fourHundred == 0 {
		return true
	} else if year % hundred == 0 {
		return false
	} else if year % four == 0 {
		return true
	}
	return false
}
