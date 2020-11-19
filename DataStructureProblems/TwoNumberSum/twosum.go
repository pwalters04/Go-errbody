package TwoNumberSum

func TwoNumberSum(array []int, target int) []int {
	numberCheck := make(map[int]bool)
	var ans []int
	var y int
	for i:=0; i <= len(array)-1 ; i++{
		y = target - (array[i])
		if numberCheck[y] == false{
			numberCheck[array[i]]=true
		}else{
			ans = append(ans,array[i])
			ans = append(ans,y)
			return ans
		}
	}
	return []int{}
}
