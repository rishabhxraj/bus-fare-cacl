package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rishabh053/bus-fare-calc/routes"
	"github.com/rishabh053/bus-fare-calc/util"
)

var Stops routes.Stops

func Prompt() (int, int) {
	Stops = routes.Load()
	var str string
	r := bufio.NewReader(os.Stdin)
	for {
		printStops(Stops)
		str, _ = r.ReadString('\n')
		if str != "" {
			break
		}
	}
	return parse(strings.TrimSpace(str))
}

func parse(input string) (int, int) {
	i := strings.Split(input, ">")
	e := util.ValidateID(toInt(i[0]), toInt(i[1]))
	util.HandleErr(e)
	return toInt(i[0]), toInt(i[1])
}

func toInt(value string) int {
	i, e := strconv.Atoi(value)
	util.HandleErr(e)
	return i
}

func printStops(stops routes.Stops) {
	fmt.Println("Route Info\tStops\tDistance\t")
	for _, stop := range stops.Location {
		data := fmt.Sprintf("%d\t%s\t%d\t", stop.ID, stop.Stop, stop.Distance)
		fmt.Println(data)
	}

}
