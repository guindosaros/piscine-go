package piscine

func SortWordArr(array []string) {
	ech:=""
	for i:=0;i<len(array);i++{
		for j:=i+1;j<len(array);j++{
			a:=[]byte(array[j])
			b:=[]byte(array[i])
			if comparearrayascii(a, b){
				ech=array[i]
				array[i]=array[j]
				array[j]=ech
			}
		}
	}
}

func comparearrayascii(a, b []byte)bool{
	n:=0
	if len(a)<=len(b){
		n=len(a)	
	}else{
		n=len(b)
	}
	for i:=0;i<n;i++{
		if b[i]<a[i]{
			return false
		}
	}
	return true
}
