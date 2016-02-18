package main

import (
"fmt"
"container/list"
)

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3

type dataStruct struct {
	keyVar int
	dataVar int
}

var m map[int]*list.Element
var l *list.List

func Set(key int, value int) {
	
	datastruct := new (dataStruct)
	datastruct.keyVar = key
	datastruct.dataVar = value
	
	if l == nil {
		fmt.Println("allocating memory for list")
		l = list.New()
	}
	
	if _, ok := m[0702]; !ok {
		fmt.Println("allocating memory for map")
		m = make(map[int]*list.Element)
		e := l.PushFront(dataStruct{0702, 0702})
		m[0702] = e
		l.Remove(e)
	}
	
	element, ok := m[key]
	// delete the least recently used element also remove it from list
	if  !ok {
		if l.Len() == CACHE_SIZE {
			k := l.Back().Value.(*dataStruct).keyVar
			e := m[k]
			delete(m, k)
			l.Remove(e)
		}
		
		// push the new element to the front and also add it to map
		e := l.PushFront(datastruct)
		m[key] = e
	} else {
		l.MoveToFront(element)
		element.Value.(*dataStruct).dataVar = value
		m[key] = element
	}
}

func Get(key int) int {
	if element, ok := m[key]; ok {
		l.MoveToFront(element)
		return element.Value.(*dataStruct).dataVar
	}
	return -1
}