// Program to illustrate the usage of
// SwapPointer function in Golang
// Including main package
package main

// Importing fmt and sync/atomic
import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

// Defining a struct type L
type L struct{ x, y, z int }

// Declaring pointer to L struct type
var PL *L

var pGet unsafe.Pointer
var pSet *L

var px, py L

var unsafepL = (*unsafe.Pointer)(unsafe.Pointer(&PL))

func initSwap() {
	// Defining *addr unsafe.Pointer

	// Storing value to the pointer
	atomic.StorePointer(
		unsafepL, unsafe.Pointer(&px))

	px.x = 1000
	px.y = 2000
	px.z = 3000

	pGet = atomic.SwapPointer(unsafepL,
		unsafe.Pointer(&py))

	pSet = &py
}

func getStruct() *L {
	return (*L)(pGet)
}

func setStruct(){
	pSet.x = (*L)(pGet).x + 1
	pSet.y = (*L)(pGet).y + 1
	pSet.z = (*L)(pGet).z + 1
	pSet   = (*L)(pGet)
	pGet = atomic.SwapPointer(unsafepL,pGet)
}

func getContinues(mtype int32) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d,%d,%d\n", getStruct().x, getStruct().y, getStruct().z)
		time.Sleep(time.Duration(mtype) * time.Millisecond)
	}
}

func setContinues()  {
	for i := 0; i < 100; i++ {
		setStruct()
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

// Main function
func main() {
	initSwap()
	go getContinues(100)
	go getContinues(200)
	go getContinues(300)
	go getContinues(400)
	go getContinues(500)
	go setContinues()
	go getContinues(600)
	go getContinues(700)
	go getContinues(800)
	go getContinues(900)
	go getContinues(1000)
	time.Sleep(time.Duration(1) * time.Hour)
}
