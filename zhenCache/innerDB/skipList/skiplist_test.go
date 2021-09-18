package skiplist

import (
	keystruct "basic/zhenCache/innerDB/keystruct"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//Implement of SlistKey
type IntTestKey struct {
	keystruct.DefaultKey
	key int
}

func (key IntTestKey) CompareBiggerThan(other keystruct.KeyStruct) bool {
	return key.key > other.KeyInt32()
}

func (key IntTestKey) KeyInt32() int {
	return key.key
}

var wg sync.WaitGroup
var skList SkipList

const MAX_OPERATION = 20000 //the max number of single CRUD operation
const LEVEL_COUNT = 32      //the level number of SkipList
const TOP_COUNT = 20030

func Test_ConcurrenyOperation(t *testing.T) {
	skList = New(LEVEL_COUNT)
	rand.Seed(time.Now().UnixNano())
	//keySlice := make([]IntTestKey, 200000)
	wg.Add(4)
	go func() {
		for i := 0; i < MAX_OPERATION; i++ {
			key := IntTestKey{keystruct.DefaultKey{}, int(rand.Uint32())}
			skList.InsertElement(key, "fuck you")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < MAX_OPERATION; i++ {
			key := IntTestKey{keystruct.DefaultKey{}, int(rand.Uint32())}
			//keySlice[i] = key
			skList.UpdateDuplicateKey(key, "fuck you")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < MAX_OPERATION; i++ {
			key := IntTestKey{keystruct.DefaultKey{}, int(rand.Uint32())}
			skList.Search(key)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < MAX_OPERATION; i++ {
			key := IntTestKey{keystruct.DefaultKey{}, int(rand.Uint32())}
			skList.Delete(key)
		}
		wg.Done()
	}()

	wg.Wait()
}

func Test_BasciFunctionTest(t *testing.T) {
	skList = New(LEVEL_COUNT)
	rand.Seed(time.Now().UnixNano())
	keySlice := make([]IntTestKey, MAX_OPERATION)
	for i := 0; i < MAX_OPERATION; i++ {
		key := IntTestKey{keystruct.DefaultKey{}, int(rand.Uint32())}
		skList.InsertElement(key, "fuck you")
	}
	for i := 0; i < MAX_OPERATION; i++ {
		key := IntTestKey{keystruct.DefaultKey{}, int(rand.Uint32())}
		keySlice[i] = key
		skList.UpdateDuplicateKey(key, "fuck you")
	}
	for i := 0; i < TOP_COUNT; i++ {
		skList.TopN(i)
	}
	for i := 0; i < MAX_OPERATION; i++ {
		key := IntTestKey{keystruct.DefaultKey{}, int(rand.Uint32())}
		skList.Search(key)
	}
	for i := 0; i < MAX_OPERATION; i++ {
		skList.Delete(keySlice[i])
	}
	skList.Show()
}