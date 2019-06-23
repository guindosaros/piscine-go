package piscine

func Unmatch(arr []int) int {
	result:=4
	for i:=range arr{
		if !ispair(arr[i],i, arr){
			result=arr[i]
			return result		
		}
	}
	return result
}

func ispair(elt,index int, arr[]int) bool{
	for i:=range arr{
		if i!=index{
			if arr[i]==elt{
				return true			
			}
		}
	}
	return false	
}