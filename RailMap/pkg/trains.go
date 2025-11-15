package pkg

import "fmt"

type Train struct {
	trainName string
	trainNum  int
}

type TrainDetails struct {
	Trains []Train
}

func (trn *Train) AddTrain(name string, num int) {

	trn.trainName = name
	trn.trainNum = num
}

func (Tr TrainDetails) GetTrain() {
	fmt.Println(Tr.Trains)
}
