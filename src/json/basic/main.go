package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type Timeout struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Ip      string `json:"ip"`
	Port    int    `json:"port"`
}

func main() {
	/* bool */
	fmt.Println("__________ bool __________")
	boolByteArray, boolError := json.Marshal(true)
	if boolError != nil {
		fmt.Println("bool error: ", boolError)
	} else {
		fmt.Println(boolByteArray)
		fmt.Println(string(boolByteArray))
	}

	/* int */
	fmt.Println("__________ int __________")
	intByteArray, intError := json.Marshal(10)
	if intError != nil {
		fmt.Println("int error: ", intError)
	} else {
		fmt.Println(intByteArray)
		fmt.Println(string(intByteArray))
	}

	/* float */
	fmt.Println("__________ float __________")
	floatByteArray, floatError := json.Marshal(1.23)
	if floatError != nil {
		fmt.Println("float error: ", floatError)
	} else {
		fmt.Println(floatByteArray)
		fmt.Println(string(floatByteArray))
	}

	/* string */
	fmt.Println("__________ string __________")
	stringData := []string{"hi", "hello", "good"}
	stringByteArray, stringError := json.Marshal(stringData)
	if stringError != nil {
		fmt.Println("string error: ", stringError)
	} else {
		fmt.Println(stringByteArray)
		fmt.Println(string(stringByteArray))
	}

	/* map int */
	fmt.Println("__________ map int __________")
	mapIntData := map[string]int{"hi": 2, "hello": 5}
	mapIntByteArray, mapIntError := json.Marshal(mapIntData)
	if mapIntError != nil {
		fmt.Println("mapInt error: ", mapIntError)
	} else {
		fmt.Println(mapIntByteArray)
		fmt.Println(string(mapIntByteArray))
	}

	/* address struct */
	/* packing */
	fmt.Println("__________ address struct __________")
	addressStructData_1 := &Address{
		Ip:   "127.0.0.1",
		Port: 1234,
	}
	addressStructByteArray_1, addressStructError_1 := json.Marshal(addressStructData_1)
	var jsonPackingAddressStruct_1 string
	if addressStructError_1 != nil {
		fmt.Println("addressStruct error: ", addressStructError_1)
		return
	} else {
		fmt.Println(addressStructByteArray_1)
		jsonPackingAddressStruct_1 = string(addressStructByteArray_1)
		fmt.Println(jsonPackingAddressStruct_1)
	}
	/* unpacking */
	var jsonUnpackingAddressStructData_1 Address
	jsonUnpackingAddressStructError_1 := json.Unmarshal(addressStructByteArray_1,
		&jsonUnpackingAddressStructData_1)
	if jsonUnpackingAddressStructError_1 != nil {
		fmt.Println("json unpacking addressStruct error: ", jsonUnpackingAddressStructError_1)
	} else {
		fmt.Printf("%+v\n", jsonUnpackingAddressStructData_1)
		fmt.Println(jsonUnpackingAddressStructData_1.Ip)
		fmt.Println(jsonUnpackingAddressStructData_1.Port)
	}

	/* address struct */
	/* packing */
	fmt.Println("__________ address struct __________")
	addressStructData_2 := &Address{
		Ip:   "127.0.0.1",
		Port: 1234,
	}
	addressStructByteArray_2, addressStructError_2 := json.Marshal(addressStructData_2)
	var jsonPackingAddressStruct_2 string
	if addressStructError_2 != nil {
		fmt.Println("addressStruct error: ", addressStructError_2)
		return
	} else {
		fmt.Println(addressStructByteArray_2)
		jsonPackingAddressStruct_2 = string(addressStructByteArray_2)
		fmt.Println(jsonPackingAddressStruct_2)
	}
	/* unpacking */
	var jsonUnpackingAddressStructData_2 map[string]interface{}
	jsonUnpackingAddressStructError_2 := json.Unmarshal(addressStructByteArray_2,
		&jsonUnpackingAddressStructData_2)
	if jsonUnpackingAddressStructError_2 != nil {
		fmt.Println("json unpacking addressStruct error: ", jsonUnpackingAddressStructError_2)
	} else {
		fmt.Printf("%+v\n", jsonUnpackingAddressStructData_2)
		fmt.Println(jsonUnpackingAddressStructData_2["ip"])
		fmt.Println(jsonUnpackingAddressStructData_2["port"])
	}
}
