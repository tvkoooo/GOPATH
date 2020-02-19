package main

import (
	"fmt"
	"xcbbrobot/common/typechange"
)

func main() {
	uid := 10005129
	fmt.Println(uid)
	struid := string(uid)
	fmt.Println(struid)
	var strsid string
	strsid = "nimeide"
	fmt.Println(strsid)

	var stringUid string
	uidInt := int(uid)
	typechange.Int2String(&uidInt, &stringUid)
	fmt.Println(stringUid)

	fmt.Println(typechange.Int2StringRe(uidInt))

}
