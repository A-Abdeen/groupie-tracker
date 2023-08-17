package gt

import "strconv"
func Arrangestring(array []string)[]string{
	var Member []string
	for i := 0;i <len(array);i++{
		
		Membermiddles := "(" + strconv.Itoa(i) + ") " + array[i] + "\n"
		Member = append(Member, Membermiddles)
		
	}	
	array = Member
	// fmt.Println(array)
	return array
} 