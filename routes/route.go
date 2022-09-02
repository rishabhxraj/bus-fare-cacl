package routes

import (
	"github.com/rishabh053/bus-fare-calc/util"

	"encoding/json"
	"io/ioutil"
	"os"
)

type Stop struct {
	ID       int    `json:"id"`
	Stop     string `json:"stop"`
	Distance int    `json:"distance"`
}

type Stops struct {
	Location []Stop `json:"stops"`
}

func Load() Stops {
	var l Stops

	// Open our jsonFile
	jsonFile, e := os.Open("routes.json")
	util.HandleErr(e)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, e := ioutil.ReadAll(jsonFile)
	util.HandleErr(e)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'l' which we defined above
	util.HandleErr(json.Unmarshal(byteValue, &l))

	return l
}
