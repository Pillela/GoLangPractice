package pkg

import (
	"fmt"
	"time"
)

var I int = 0

var hour_message, minute_message, second_message string = "Bong", "Tock", "Tick"

type Clockdata struct{
	Index int
	Message string
}

var Dataarray[10801] Clockdata

func TickTockBong(){
	go func(second_message *string){
		var replaced_string string
		fmt.Scanln(&replaced_string)
		*second_message = replaced_string
	}(&second_message)

	for {
			I++
			//time.Sleep(1 * time.Second)
			if I%int(3*60*60)==0{
				Dataarray[I].Index = I
				Dataarray[I].Message = hour_message
				fmt.Println(Dataarray[I].Message, I)
				return
			}else if I%int(60*60)==0{
				Dataarray[I].Index = I
				Dataarray[I].Message = hour_message
				fmt.Println(Dataarray[I].Message, I)
			}else if I%int(60)==0{
				Dataarray[I].Index = I
				Dataarray[I].Message = minute_message
				fmt.Println(Dataarray[I].Message, I)
			}else{
				Dataarray[I].Index = I
				Dataarray[I].Message = second_message
				fmt.Println(Dataarray[I].Message, I)
			}
	}
}
