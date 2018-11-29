package debug_loops

/*

func main() {

	traps_list := NewTrapsList()

	for i := 0; i < 20; i++ {
		traps_list.Trap("main", 10, 15)
	}

}

*/

import (
	"fmt"
)

type Trap struct {
	Name string
	Value uint
	WarningThreshold uint
	PanicThreshold uint
}

type TrapsList map[string]*Trap

func NewTrapsList() TrapsList {
	list := make(TrapsList, 1)
	return list
}

func (list TrapsList) Trap(name string, warning uint, panic uint) {
	var trap *Trap
	if _, ok := list[name]; !ok {
		trap = &Trap{Name: name, Value: 0, WarningThreshold: warning, PanicThreshold: panic}
		list[name] = trap
	} else {
		trap = list[name]
	}

	trap.Iterate()
}

func (trap *Trap) Iterate() {
	trap.Value ++
	if trap.Value > trap.WarningThreshold {
		fmt.Printf("Warning: %s = %d\n", trap.Name, trap.Value)
	}

	if trap.Value > trap.PanicThreshold {
		panic(fmt.Sprintf("Problem here: %s = %d\n", trap.Name, trap.Value))
	}
}



