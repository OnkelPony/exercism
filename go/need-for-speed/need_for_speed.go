package speed

// Car represents vehicle with fields batter, batteryDrain, speed and distance.
type Car struct {
	battery      int
	batteryDrain int
	distance     int
	speed        int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		distance:     0,
		speed:        speed}
}

// Track represents track with field distance.
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	virtBatt := car.battery - car.batteryDrain
	if virtBatt > 0 {
		car.distance = car.speed
		car.battery = virtBatt
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return track.distance/car.speed*car.batteryDrain <= car.battery
}
