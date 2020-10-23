package main 
import ("fmt")

func concat(slices [][]int) []int{
	/*
	in: slice of integer slices
	out: single slice of all slicles concatenated
	*/
	ans := slices[0][:]
	if len(slices)>1{
		for i:=1; i<len(slices); i++{
			ans = append(ans, slices[i]...)
		}
	}
	return ans
	
}

func tmp(){
	left := []int{1,2,3}
	center := []int{4,5,6}
	right := []int{7,8,9}
	group := make([][]int, 3)
	group[0] = left
	group[1] = center
	group[2] = right
	ans := concat(group)
	fmt.Println(ans)

}

func quickSort(array []int) []int{
	var left []int
	var center []int
	var right []int

	if len(array)>1{
		pivot := array[0]
		for i:=0; i<len(array); i++{
			if array[i]<pivot{
				left = append(left, array[i])
			} else if array[i]==pivot{
				center = append(center, array[i])
			} else{
				right = append(right, array[i])
			}
		}
		tmp := make([][]int, 3)
		tmp[0] = quickSort(left)
		tmp[1] = center
		tmp[2] = quickSort(right)
		return concat(tmp)
	} else{
		return array
	}
}


func main(){
	array := []int{2,6,1,15,3,40,8,10,9,50}
	ans := quickSort(array)
	fmt.Println("Números ordeandos: ", ans)
	//tmp()
}