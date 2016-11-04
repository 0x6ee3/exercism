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

func addComputeCell(c Cell, cc *computeCell) {
	if cx, ok := c.(*computeCell); ok {
		cx.addComputeCell(cc)
	} else if ic, ok := c.(*inputCell); ok {
		ic.addComputeCell(cc)
	}
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
		cc.SetValue(newValue)

		for _, cb := range cc.callbacks {
			cb(newValue)
		}
	}
}

func (cc *computeCell) AddCallback(f func(int)) CallbackHandle {
	cc.callbacks[&f] = f
	return &f
}

func (cc *computeCell) RemoveCallback(cb CallbackHandle) {
	delete(cc.callbacks, cb)
}

// -- Reactor Methods ---------------------------------------------------------

// New ..
func New() Reactor {
	return &CellReactor{}
}

// CellReactor ..
type CellReactor struct{}

// CreateInput ..
func (r *CellReactor) CreateInput(value int) InputCell {
	return &inputCell{value: value}
}

// CreateCompute1 ..
func (r *CellReactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	cc := computeCell{
		callbacks: make(map[CallbackHandle]func(int)),
		calcValue: func() int { return f(c.Value()) },
	}
	cc.value = f(c.Value())
	addComputeCell(c, &cc)
	return &cc
}

// CreateCompute2 ..
func (r *CellReactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	cc := computeCell{
		callbacks: make(map[CallbackHandle]func(int)),
		calcValue: func() int { return f(c1.Value(), c2.Value()) },
	}
	cc.value = f(c1.Value(), c2.Value())
	addComputeCell(c1, &cc)
	addComputeCell(c2, &cc)
	return &cc
}
