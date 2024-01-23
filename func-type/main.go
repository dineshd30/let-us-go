package main

import (
	"fmt"
	"strings"
	"time"
)

// Create function type Message
type Message func(name string) string

// Create toLower receiver function
func (m Message) toLower(name string) string {
	return strings.ToLower(m(name))
}

// Create toUpper receiver function
func (m Message) toUpper(name string) string {
	return strings.ToUpper(m(name))
}

func main() {
	// Create instance of Message function by passing function literal
	msg := Message(func(name string) string {
		return "Hey, " + name
	})

	fmt.Println(msg("Dinesh"))
	fmt.Println(msg.toUpper("Dinesh"))
	fmt.Println(msg.toLower("Dinesh"))
	// minMilli := int64(0)
	// minTime := time.UnixMilli(0)
	// maxMilli := int64(253402214400000) //  9999-12-31 00:00:00 +0000 UTC
	// maxTime := time.UnixMilli(maxMilli)
	// nowMilli := time.Now().UTC().UnixMilli()
	// nowTime := time.UnixMilli(nowMilli)
	// fmt.Printf("%v, %v, %v \n", minMilli, maxMilli, nowMilli)
	// fmt.Printf("%v, %v, %v \n", minTime, maxTime, nowTime)
	// fmt.Println(nowTime.After(minTime))
	// fmt.Println(nowTime.Before(maxTime))

	mt := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	fmt.Println(mt)
	mtm := mt.UnixMilli()
	fmt.Println(mtm)
	mtmog := time.UnixMilli(mtm)
	fmt.Println(mtmog)

	mt1 := time.Date(1970, 01, 01, 0, 0, 0, 0, time.UTC)
	fmt.Println(mt1)
	mtm1 := mt1.UnixMilli()
	fmt.Println(mtm1)
	mtmog1 := time.UnixMilli(mtm1)
	fmt.Println(mtmog1)
}
