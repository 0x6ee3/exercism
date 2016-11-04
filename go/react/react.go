package react

const testVersion = 4

// -- Cell Types --------------------------------------------------------------

type inputCell struct {
	value        int
	computeCells []*computeCell
}

func (i *inputCell) Value() int {
	return i.value
}

func (i *inputCell) SetValue(v int) {
	i.value = v

	for _, c := range i.computeCells {
		c.Update()
	}
}

func (i *inputCell) addComputeCell(cc *computeCell) {
	i.computeCells = append(i.computeCells, cc)
}

type computeCell struct {
	inputCell
	calcValue func() int
	callbacks map[CallbackHandle]func(int)
}

func (cc *computeCell) Value() int {
	return cc.calcValue()
}

func (cc *computeCell) Update() {
	if newValue := cc.Value(); newValue != cc.value {
		cc.value = newValue

		for _, c := range cc.computeCells {
			c.Update()
		}

		for _, cb := range cc.callbacks {
			cb(newValue)
		}
	}
}

func (cc *computeCell) AddCallback(f func(int)) CallbackHandle {
	cbh := &f
	cc.callbacks[cbh] = f
	return cbh
}

func (cc *computeCell) RemoveCallback(cb CallbackHandle) {
	delete(cc.callbacks, cb)
}

// -- Reactor Methods ---------------------------------------------------------

// CellReactor ..
type CellReactor struct{}

// CreateInput ..
func (r *CellReactor) CreateInput(value int) InputCell {
	return &inputCell{value: value}
}

// CreateCompute1 ..
func (r *CellReactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	cc := computeCell{callbacks: make(map[CallbackHandle]func(int))}

	if cx, ok := c.(*computeCell); ok {
		cx.addComputeCell(&cc)
	} else if ic, ok := c.(*inputCell); ok {
		ic.addComputeCell(&cc)
	}

	cc.value = f(c.Value())
	cc.calcValue = func() int {
		return f(c.Value())
	}

	return &cc
}

// CreateCompute2 ..
func (r *CellReactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	cc := computeCell{callbacks: make(map[CallbackHandle]func(int))}

	if cx, ok := c1.(*computeCell); ok {
		cx.addComputeCell(&cc)
	} else if ic, ok := c1.(*inputCell); ok {
		ic.addComputeCell(&cc)
	}
	if cx, ok := c2.(*computeCell); ok {
		cx.addComputeCell(&cc)
	} else if ic, ok := c2.(*inputCell); ok {
		ic.addComputeCell(&cc)
	}

	cc.value = f(c1.Value(), c2.Value())
	cc.calcValue = func() int {
		return f(c1.Value(), c2.Value())
	}

	return &cc
}

// -- New ---------------------------------------------------------------------

// New ..
func New() Reactor {
	return &CellReactor{}
}
