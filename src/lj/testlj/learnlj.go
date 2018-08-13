package main

import "fmt"


type ljljteststr1 struct {
	h int
	l int
	d int
}

type ljljteststr2 struct {
	mm []ljljteststr1
}

type ljljteststr3 struct {
	xx int
	yy ljljteststr2
}


func main() {
	var ljlj ljljteststr3
	var ljljhav ljljteststr3
	ljlj.xx = 55

	fmt.Println(" ljlj:",ljlj)

	ljljhav.xx =9999

	var kkkk1 ljljteststr1
	kkkk1.d = 4
	kkkk1.h = 5
	kkkk1.l = 3
	fmt.Println(" kkkk1:",kkkk1)


	ljljhav.yy.mm =  make([]ljljteststr1,3)
	ljljhav.yy.mm[0] = kkkk1
	ljljhav.yy.mm[1] = kkkk1
	ljljhav.yy.mm[2] = kkkk1

	fmt.Println(" ljljhav:",ljljhav)

	if 0!= len(ljlj.yy.mm){
		fmt.Println(" ljlj has num:",ljlj)
	}
	if 0!= len(ljljhav.yy.mm){
		fmt.Println(" ljljhav has num:",ljljhav)
	}

	for _,v := range ljljhav.yy.mm{
		if 3 == v.l{
			fmt.Println(" ljljhav: has 3")
			break
		}
	}
	fmt.Println(" over")




}