package syncmap

import (
	"fmt"
	"sync"
	"testing"
)

type SkillTable struct {
	Id int
}

func TestSyncMap(t *testing.T) {
	confs := sync.Map{}
	confs.Store("skill_table", &SkillTable{Id: 1})
	value, ok := confs.Load("skill_table")
	if !ok {
		return
	}
	skillTable := value.(*SkillTable)
	fmt.Println(skillTable.Id)
	confs.Load("skill_table")

	confs.LoadOrStore("skill_table2", &SkillTable{Id: 2})

	confs.Range(func(key, value any) bool {
		return false
	})

	_, loaded := confs.LoadAndDelete("skill_table3")
	if !loaded {
		return
	}

	val, loaded := confs.LoadAndDelete("skill_table2")
	if !loaded {
		return
	}
	table2 := val.(*SkillTable)
	fmt.Println(table2.Id)

}
