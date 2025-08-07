package say

import "fmt"

type SayToConsole struct {}

func NewSayToConsole() *SayToConsole{
	return &SayToConsole{}
}

func (s *SayToConsole) Log(args... any) {
	fmt.Println(args...)
}
