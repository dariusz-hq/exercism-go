package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
    for i := 0; i < len(birdsPerDay); i++ {
        total += birdsPerDay[i]
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    if week < 1 {
        panic("Week cannot be less than 1")
    }
	total := 0
    daysInWeek := 7
    init := (daysInWeek * week) - 7
    max := (week * daysInWeek) + 1
    if len(birdsPerDay) < max {
        max = len(birdsPerDay)
    }
    for i := init; i < max; i++ {
        total = total + birdsPerDay[i]
    }
    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
        if i % 2 == 0 {
            birdsPerDay[i]++
        }
    }
    return birdsPerDay
}
