package main

import (
	"fmt"
	"maps"
)

type StringIntMap struct {
	Storage map[string]int
}

func main() {

	mapa := StringIntMap{Storage: make(map[string]int)}
	fmt.Println(mapa)

	mapa.Add("first", 1)
	mapa.Add("second", 2)
	mapa.Add("third", 3)
	fmt.Println(mapa)

	mapa.Remove("first")
	fmt.Println(mapa)

	second_map := mapa.Copy()
	fmt.Println(second_map)

	Is := mapa.Exists("second")
	fmt.Println(Is)
	Is = mapa.Exists("forth")
	fmt.Println(Is)

	v, Is := mapa.Get("third")
	fmt.Println(v, Is)

}

func (m *StringIntMap) Add(key string, value int) {
	m.Storage[key] = value

}

func (m *StringIntMap) Remove(key string) {

	delete(m.Storage, key)

}

func (m *StringIntMap) Copy() map[string]int {

	stor := make(map[string]int)

	maps.Copy(stor, m.Storage)

	return stor

}

func (m *StringIntMap) Exists(key string) bool {

	_, ok := m.Storage[key]

	return ok

}

func (m *StringIntMap) Get(key string) (int, bool) {

	v, ok := m.Storage[key]

	return v, ok

}
