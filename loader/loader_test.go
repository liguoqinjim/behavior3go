package loader

import (
	"fmt"
	"reflect"
	"testing"

	//b3 "behavior3go"
	//. "behavior3go/actions"
	//. "behavior3go/composites"
	. "behavior3go/config"
	. "behavior3go/core"
	//. "behavior3go/decorators"
)

type Test struct {
	value string
}

func (test *Test) Print() {
	fmt.Println(test.value)
}

func TestExample(t *testing.T) {
	maps := createBaseStructMaps()
	if data, err := maps.New("Runner"); err != nil {
		t.Error("Error:", err, data)
	} else {
		t.Log(reflect.TypeOf(data))
	}

}

func TestLoadTree(t *testing.T) {
	treeConfig, ok := LoadTreeCfg("tree.json")
	if ok {
		tree := CreateBevTreeFromConfig(treeConfig, nil)
		tree.Print()

		board := NewBlackboard()
		for i := 0; i < 5; i++ {
			tree.Tick(i, board)
		}
	} else {
		t.Error("LoadTreeCfg err")
	}

}
