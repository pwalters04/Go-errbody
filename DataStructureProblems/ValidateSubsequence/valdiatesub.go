package ValidateSubsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	var checkArrayLength int
	var sequenceLength= len(sequence)
	var arrayLength = len(array)
	var checkSequenceLength int

	matchCount:=0
	j:=0
	i:=0

	for {
		if sequence[i] == array[j]{
			checkSequenceLength++
			i++
			matchCount++
		}
		checkArrayLength++
		j++
		if matchCount == sequenceLength{
			return true
		}

		if checkArrayLength == arrayLength{
			break
		}
	}
	return false
}
