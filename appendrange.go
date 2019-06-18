package piscine

func AppendRange(min, max int) []int {
        if min >= max {
                var tab []int = nil
                return tab
        }else{
                tab := []int{}
		
                for i:= min;i < max;i++ {
                        tab = append(tab,i)
                }
                return tab
        }
}
