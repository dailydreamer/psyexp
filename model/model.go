package model

import (
	"fmt"
	"time"
	"container/list"
	"os"
	"log"
	"encoding/csv"
	"math/rand"
	"io/ioutil"
	"math/big"
	"github.com/Ramshackle-Jamathon/go-quickPerm"
)

type Tester struct {
	ID string
	StartTime time.Time
	DecisionTime time.Duration
	CurrentPicture *list.Element
	PicturesList *list.List
	PicturePicked string
}

type Picture struct {
	PictureNames []string
	PictureNumsFac big.Int
}

var pic Picture

func Init() {
	files, err := ioutil.ReadDir("static/images/")
	if err != nil {
		log.Fatal(err)
	}
	num := len(files)
	pic.PictureNumsFac.MulRange(1, int64(num))
	for _, file := range files {
		pic.PictureNames = append(pic.PictureNames, file.Name())
	}
	log.Println(pic)
}

func New() *Tester {
	t := &Tester{}
	t.PicturesList = list.New()

	idx := rand.Int63n(pic.PictureNumsFac.Int64())
	for permutation := range quickPerm.GeneratePermutationsString(pic.PictureNames) {
		if (idx == 0) {
			for _, item := range permutation {
				t.PicturesList.PushBack(item)
			}
			break
		}
		idx = idx - 1
	}

	// for i := 1; i <= 10; i++ {
	// 	t.PicturesList.PushBack(fmt.Sprintf("%d",i))
	// }
	return t
}

func (t *Tester) Start() interface{} {
	t.StartTime = time.Now()
	t.CurrentPicture = t.PicturesList.Front()
	return t.CurrentPicture.Value
}

func (t *Tester) Finish() interface{} {
	t.PicturePicked = t.CurrentPicture.Value.(string)
	t.DecisionTime = time.Since(t.StartTime)
	t.saveToCsv()
	return t.PicturePicked
}

func (t *Tester) saveToCsv() {
	data := []string{
		t.ID, 
		t.PicturePicked, 
		fmt.Sprintf("%f", t.DecisionTime.Seconds()),
	}
	fileName := "psyexp.csv"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("\nOpen csv file error: %s\n", err.Error())
		for i := 2; err != nil; i++ {
			file, err = os.OpenFile(fmt.Sprintf("psyexp%d.csv",i), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		}
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(data)
	if err != nil {
		log.Printf("\nCannot write to csv: %s\n", err.Error())
	}
	log.Println("exp data:")
	log.Println(data)
}

// Keep return next pid, isRoundOver
func (t *Tester) Keep() (interface{}, bool) {
	isRoundOver := false
	if t.CurrentPicture.Next() == nil {
		// round over
		isRoundOver = true
		t.CurrentPicture = t.PicturesList.Front()
	} else {
		t.CurrentPicture = t.CurrentPicture.Next()
	}
	return t.CurrentPicture.Value, isRoundOver
}

// Giveup return next pid, isAllOver, isRoundOver
func (t *Tester) Giveup() (interface{}, bool, bool) {
	eleNeedToBeDelete := t.CurrentPicture
	isRoundOver := false
	if t.CurrentPicture.Next() == nil {
		// round over
		isRoundOver = true
		t.CurrentPicture = t.PicturesList.Front()
	} else {
		t.CurrentPicture = t.CurrentPicture.Next()
	}
	t.PicturesList.Remove(eleNeedToBeDelete)
	if t.CurrentPicture.Next() == nil && t.CurrentPicture.Prev() == nil {
		// only left one picture, which is the choosen one
		return t.CurrentPicture.Value, true, isRoundOver
	} else {
		return t.CurrentPicture.Value, false, isRoundOver
	}
}