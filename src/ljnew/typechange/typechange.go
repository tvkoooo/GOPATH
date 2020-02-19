package typechange

import "strconv"

func String2Int(stringIn *string, intOut *int) (ferr error) {
	*intOut, ferr = strconv.Atoi(*stringIn)
	return ferr
}

func String2Int64(stringIn *string, intOut *int64) (ferr error) {
	*intOut, ferr = strconv.ParseInt((*stringIn), 10, 64)
	return ferr
}

func Int2String(intIn *int, stringOut *string) {
	*stringOut = strconv.Itoa(*intIn)
}

func Int642String(intIn *int64, stringOut *string) {
	*stringOut = strconv.FormatInt(*intIn, 10)
}

func Slice2String(sliceIn *[]byte, stringOut *string) {
	*stringOut = string((*sliceIn)[:])
}

func String2Slice(stringIn *string, sliceOut *[]byte) {
	*sliceOut = []byte(*stringIn)
}
