package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	no := 100
	fmt.Println(reflect.TypeOf(no))

	intstr := "100"
	fourBaseEightBitInt, _ := strconv.ParseInt(intstr, 4, 8)
	tenBaseSixteenBitInt, _ := strconv.ParseInt(intstr, 10, 16)

	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))

	noPlayer := 8
	strPlayer := strconv.Itoa(noPlayer)
	fmt.Println(reflect.TypeOf(strPlayer))
}
