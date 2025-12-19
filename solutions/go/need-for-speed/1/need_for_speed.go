package speed

// Car represents type
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
        distance: 0,
    }
}

// Track type represents racing track
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if (car.battery - car.batteryDrain) < 0 {
        return car
    }
    
	car.distance += car.speed
    car.battery -= car.batteryDrain
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	batteryCycles := car.battery / car.batteryDrain
    possibleDistance := batteryCycles * car.speed

    if float64(track.distance) <= float64(possibleDistance) {
        return true
    }

    return false
}
