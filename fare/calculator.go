package fare

import (
	"fmt"

	"github.com/rishabh053/bus-fare-calc/cmd"
	"github.com/rishabh053/bus-fare-calc/routes"
)

func Start() {
	srcID, destID := cmd.Prompt()
	src, dest := GetDistance(srcID, destID)
	fare := CalculateFare(dest.Distance - src.Distance)
	out := fmt.Sprintf("%s > %s = Rs.%d\n", src.Stop, dest.Stop, fare)
	fmt.Printf(out)
}

func GetDistance(src, dest int) (routes.Stop, routes.Stop) {
	stops := cmd.Stops
	source := get(stops, src)
	destination := get(stops, dest)
	return source, destination
}

func get(stops routes.Stops, id int) routes.Stop {
	for _, stop := range stops.Location {
		if stop.ID == id {
			return stop
		}
	}
	return routes.Stop{}
}

func CalculateFare(distance int) int {
	minFare := 3
	if distance <= 3 {
		return minFare
	} else if distance > 3 && distance <= 20 {
		d := distance - minFare
		return minFare + d*2
	} else if distance >= 20 {
		d := distance - 20
		return minFare + (20-3)*2 + d
	}
	return 0
}
