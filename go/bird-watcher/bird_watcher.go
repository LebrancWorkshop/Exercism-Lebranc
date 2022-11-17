package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0; 
	for _, bird := range birdsPerDay {
		total += bird; 
	}
	return total; 
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
// (week * 7 - 7) to (week * 7 - 1)
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0; 
	firstDayOfWeek := (week * 7) - 7;
	lastDayOfWeek := (week * 7) - 1;

	for index, bird := range birdsPerDay {
		if index >= firstDayOfWeek && index <= lastDayOfWeek {
			total += bird;
		}
	}

	return total; 
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, bird := range birdsPerDay {
		if index % 2 == 0 {
			birdsPerDay[index] = bird + 1; 
		}
	}
	return birdsPerDay; 
}
