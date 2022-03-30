package main

import (
	"fmt"
	"sort"
)

type Annotation struct {
	key   string
	value string
}
type Person struct {
	Name string
}

func main() {
	annotations := []Annotation{{key: "bwosa_topic", value: "true"}, {key: "adisable_woc", value: "true"}}
	sort.Slice(annotations, func(i, j int) bool {
		return annotations[i].key <= annotations[j].key
	})

	var idx int = sort.Search(len(annotations), func(i int) bool {
		return annotations[i].key >= "bwosa_topicx"

	})

	fmt.Println(idx, annotations)

	// pesquisar por sort.Search

}
