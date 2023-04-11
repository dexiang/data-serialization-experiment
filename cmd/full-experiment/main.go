package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/protobuf/proto"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/vmihailenco/msgpack/v5"

	pb "data-serialization-experiment/cmd/full-experiment/protobuf/compiled"
)

type ExperimentMatrix struct {
	Size              int
	SerializingTime   time.Duration
	DeserializingTime time.Duration
}

type User struct {
	ID      int
	Email   string
	Name    string
	IsAdmin bool
	Assets  float64
}

func generator(times int) (
	jResult ExperimentMatrix,
	mResult ExperimentMatrix,
	pResult ExperimentMatrix,
) {
	// Fake Data
	var users []User
	for i := 0; i < times; i++ {
		users = append(users, User{
			ID:      i,
			Name:    gofakeit.Name(),
			Email:   gofakeit.Email(),
			IsAdmin: gofakeit.Bool(),
			Assets:  float64(gofakeit.Float32Range(-1000*1000, 1000*1000*1000)),
		})
	}

	// ======== JSON ========
	// serialization
	jSStart := time.Now()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	jResult.Size = len(jsonBytes)
	jResult.SerializingTime = time.Since(jSStart)
	// deserialization
	jDStart := time.Now()
	var newUsersForJSON []User
	err = json.Unmarshal(jsonBytes, &newUsersForJSON)
	if err != nil {
		panic(err)
	}
	jResult.DeserializingTime = time.Since(jDStart)

	// ======== MessagePack ========
	// serialization
	mSStart := time.Now()
	msgpackBytes, err := msgpack.Marshal(users)
	if err != nil {
		panic(err)
	}
	mResult.Size = len(msgpackBytes)
	mResult.SerializingTime = time.Since(mSStart)
	// deserialization
	mDStart := time.Now()
	var newUsersForM []User
	err = msgpack.Unmarshal(msgpackBytes, &newUsersForM)
	if err != nil {
		panic(err)
	}
	mResult.DeserializingTime = time.Since(mDStart)

	// ======== Protocol Buffers ========
	// serialization
	pSStart := time.Now()
	usersPb := &pb.Users{}
	for _, user := range users {
		usersPb.User = append(usersPb.User, &pb.User{
			ID:      int32(user.ID),
			Email:   user.Email,
			Name:    user.Name,
			IsAdmin: user.IsAdmin,
			Assets:  float32(user.Assets),
		})
	}
	pbBytes, err := proto.Marshal(usersPb)
	if err != nil {
		panic(err)
	}
	pResult.Size = len(pbBytes)
	pResult.SerializingTime = time.Since(pSStart)
	// deserialization
	pDStart := time.Now()
	newUsersPb := pb.Users{}
	err = proto.Unmarshal(pbBytes, &newUsersPb)
	if err != nil {
		panic(err)
	}
	pResult.DeserializingTime = time.Since(pDStart)

	return jResult, mResult, pResult
}

func prettyByteSize(b int) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}

func main() {

	testCases := map[string]int{"Tiny": 20, "Small": 300, "Medium": 3000, "Large": 100 * 1000, "Huge": 5000 * 1000}
	//testCases := map[string]int{"Tiny": 20}
	//testCases := map[string]int{"Small": 300}
	//testCases := map[string]int{"Medium": 3000}
	//testCases := map[string]int{"Large": 100 * 1000}
	//testCases := map[string]int{"Huge": 5000 * 1000}
	keys := make([]string, 0, len(testCases))

	for k := range testCases {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	st := table.NewWriter()
	st.SetOutputMirror(os.Stdout)
	st.AppendHeader(table.Row{"#", "JSON", "MessagePack", " Protocol Buffers"})
	mt := table.NewWriter()
	mt.SetOutputMirror(os.Stdout)
	mt.AppendHeader(table.Row{"#", "JSON", "MessagePack", " Protocol Buffers"})
	ut := table.NewWriter()
	ut.SetOutputMirror(os.Stdout)
	ut.AppendHeader(table.Row{"#", "JSON", "MessagePack", " Protocol Buffers"})

	for _, k := range keys {
		jResult, mResult, pResult := generator(testCases[k])
		st.AppendRow([]interface{}{k, prettyByteSize(jResult.Size), prettyByteSize(mResult.Size), prettyByteSize(pResult.Size)})
		mt.AppendRow([]interface{}{k, jResult.SerializingTime, mResult.SerializingTime, pResult.SerializingTime})
		ut.AppendRow([]interface{}{k, jResult.DeserializingTime, mResult.DeserializingTime, pResult.DeserializingTime})
	}

	fmt.Println("Size")
	st.Render()

	fmt.Println("serializing time")
	mt.Render()

	fmt.Println("deserializing time")
	ut.Render()
}
