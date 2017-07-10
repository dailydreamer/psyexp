package model

import (
	"time"
)

type Tester struct {
	ID string
	StartTime time.Time
	DecisionTime time.Time
	CurrentPicture int
	PicturesList []int
	PicturePicked string
}

func InitTester(t *Tester) {

}

func StartTimer(t *Tester) {
	
}

func DeletePicture(t *Tester) {
	
}

func PickPicture(t *Tester) {
	
}