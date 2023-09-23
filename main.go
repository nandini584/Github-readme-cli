package main 
import (
	"fmt"
)
func main(){
	names := []string{"a","b","c"}
	id := []int{1,2,3}
	u :=make(map [string]int)
	for i:=0;i<len(names);i++{
	u[names[i]]=id[i]
}
fmt.Println(u)
}