package main
import "fmt"

func main(){
	var slice [] int;
	for i:=0;i<10;i++{
		slice = append(slice, i)
	}
	fmt.Println(slice)
}