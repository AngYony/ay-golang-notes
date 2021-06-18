// package main

// import (
// 	trippb "coolcar/proto/gen/go"
// 	"encoding/json"
// 	"fmt"

// 	"google.golang.org/protobuf/proto"
// )

// func main() {
// 	trip := trippb.Trip{
// 		Start:       "abc",
// 		End:         "def",
// 		DurationSec: 3600,
// 		FeeCent:     10000,
// 		StartPos: &trippb.Location{
// 			Latitude:  35,
// 			Longitude: 100,
// 		},
// 		EndPos: &trippb.Location{
// 			Latitude:  40,
// 			Longitude: 130,
// 		},
// 		PathLocations: []*trippb.Location{
// 			{
// 				Latitude:  50,
// 				Longitude: 100,
// 			},
// 			{
// 				Latitude:  66,
// 				Longitude: 77,
// 			},
// 		},

// 		Status: trippb.TripStatus_IN_PROGRESS,
// 	}
// 	fmt.Println(&trip)
// 	//将类型编码为二进制流
// 	b, err := proto.Marshal(&trip)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%X \n", b)

// 	var trip2 trippb.Trip
// 	//将二进制编码进行解码
// 	err = proto.Unmarshal(b, &trip2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(&trip2)

// 	b, err = json.Marshal(&trip2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", b)
// }
