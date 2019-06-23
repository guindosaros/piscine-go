package main

import "fmt"

type Door struct{
	state bool
}
const OPEN = false
const CLOSE =true

func OpenDoor(ptrDoor *Door){
	fmt.Println("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door){
	fmt.Println("Door closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor Door) bool {
	fmt.Println("Door is open ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor Door) bool {
	fmt.Println("Door is close ?")
	return ptrDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}