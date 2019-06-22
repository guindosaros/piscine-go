package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	if len(tab)>1{
		if f(tab[0],tab[1])>=0{
			for i:=0;i<len(tab)-1;i++{
				if f(tab[i],tab[i+1])<0{
					return false		
				}
			}
		}
		if f(tab[0],tab[1])<=0{
			for i:=0;i<len(tab)-1;i++{
				if f(tab[i],tab[i+1])>0{
					return false		
				}
			}
		}
	}
	
	return true
}

