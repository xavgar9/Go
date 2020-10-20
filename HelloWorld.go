package main 
  
import "fmt"
  
// Main function 
/* es es comentario también */
func solve(){
	x := [6]int{1,3,5,7,9,11}
	y := [6]int{2,4,6,8,10,12}
	var ans [12]int
	var i, j, k = 0, 0, 0
	for ok:=true; ok; {
		if i==len(x){
			for _j:=j; _j<len(y); _j++{
				ans[k] = y[_j]
				k++
			}
		} else if j==len(y){
			for _i:=i; _i<len(x); _i++{
				ans[k] = x[_i]
				k++
			}
		} else{
			if x[i]<y[j]{
				ans[k] = x[i]
				i++
			} else if y[j]<x[i]{
				ans[k] = y[j]
				j++
			} else{
				ans[k] = x[i]
				ans[k+1] = y[j]
				i++
				j++
				k++
			}
		}
		k++
		if k>=len(ans){
			ok = false
		}
	}
	fmt.Println(ans)
}


func main() { 
	/*
    var x = 3
    fmt.Println(x) 
    fmt.Printf("El tipo de dato de x es: %d \n", x)
    */
    solve()
}