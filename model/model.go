package model

import (
	"fmt"
	"time"
	"container/list"
)

type Tester struct {
	ID string
	StartTime time.Time
	DecisionTime time.Duration
	CurrentPicture *list.Element
	PicturesList *list.List
	PicturePicked string
}

func New() *Tester {
	// TODO add random
	t := &Tester{}
	t.PicturesList = list.New()
	for i := 0; i < 10; i++ {
		t.PicturesList.PushBack(fmt.Sprintf("%d",i))
	}
	return t
}

func (t *Tester) Start() interface{} {
	t.StartTime = time.Now()
	t.CurrentPicture = t.PicturesList.Front()
	return t.CurrentPicture.Value
}

func (t *Tester) Finish(pid string) interface{} {
	t.PicturePicked = pid
	t.DecisionTime = time.Since(t.StartTime)
	fmt.Printf("Decistion time: %fs\n", t.DecisionTime.Seconds())
	// TODO save to csv
	return pid
}

func (t *Tester) Keep() interface{} {
	if t.CurrentPicture.Next() == nil {
		t.CurrentPicture = t.PicturesList.Front()
	} else {
		t.CurrentPicture = t.CurrentPicture.Next()
	}
	return t.CurrentPicture.Value
}

func (t *Tester) Giveup() (interface{}, bool) {
	eleNeedToBeDelete := t.CurrentPicture
	if t.CurrentPicture.Next() == nil {
		t.CurrentPicture = t.PicturesList.Front()
	} else {
		t.CurrentPicture = t.CurrentPicture.Next()
	}
	t.PicturesList.Remove(eleNeedToBeDelete)
	if t.CurrentPicture.Next() == nil && t.CurrentPicture.Prev() == nil {
		// only left one picture, which is the choosen one
		return t.CurrentPicture.Value, true
	} else {
		return t.CurrentPicture.Value, false
	}
}

