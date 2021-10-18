package main

import (
	"fmt"
)	


type Movie struct {
	Length  int
	Name string
}
func (m *Movie) Rate(float32) {
//	if 
//	m.counter += 1
//m.counter = 0
	
}

func (m *Movie) Play(viewers int){
	return m.Playings 
}

func (m Movie) Viewers() int{
	return //return int # of people who have viewed movie
}

func (m Movie) Plays() int{
	return //return int # of times the movie has been played. total ratings divided by times movie is played
}

func (m Movie) Rating() float64{
	return //returns the rating float64 of the movie. 
}

func (m Movie) String() string{
	return Name, Length, Rate //returns name, length, and rating
}

type Theatre struct {

}

func (t *Theatre) Play(int){

}

func (t Theatre) Crtique := []string{"John","Paul","George","Ringo"{

}
func main() {
	a_Movie := Movie{6, "Endgame"}

	fmt.Println(a_Movie)
	fmt.Println(Length)
	fmt.Println("Success!")

}
