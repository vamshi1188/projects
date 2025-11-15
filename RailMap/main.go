package main

import (
	"RailMap/pkg"
	// "fmt"
)

func main() {

	var TrainData pkg.TrainDetails

	TrainData.Trains()

	TrainData.Trains("krishna", 17405)
	// fmt.Println(TrainData)
	TrainData.SetTrains("Satavahana SF Express", 12714)

	// fmt.Println(TrainData)
	TrainData.GetTrainDetails()

}
