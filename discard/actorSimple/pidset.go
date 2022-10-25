package actorNew

type PIDSet struct {
	pids   []IPid
	lookup map[int64]IPid
}

// NewPIDSet returns a new PIDSet with the given pids.
func NewPIDSet(pids ...IPid) *PIDSet {
	p := &PIDSet{}
	for _, pid := range pids {
		p.Add(pid)
	}
	return p
}

func (p *PIDSet) ensureInit() {
	if p.lookup == nil {
		p.lookup = make(map[int64]IPid)
	}
}

func (p *PIDSet) indexOf(v IPid) int {
	for i, pid := range p.pids {
		if v.Id() == pid.Id() {
			return i
		}
	}
	return -1
}

func (p *PIDSet) Contains(v IPid) bool {
	return p.lookup[v.Id()] != nil
}

// Add adds the element v to the set.
func (p *PIDSet) Add(v IPid) {
	p.ensureInit()
	if p.Contains(v) {
		return
	}
	p.lookup[v.Id()] = v
	p.pids = append(p.pids, v)
}

// Remove removes v from the set and returns true if them element existed.
func (p *PIDSet) Remove(v IPid) bool {
	p.ensureInit()
	i := p.indexOf(v)
	if i == -1 {
		return false
	}

	delete(p.lookup, v.Id())

	p.pids = append(p.pids[:i], p.pids[i+1:]...)

	return true
}

// Len returns the number of elements in the set.
func (p *PIDSet) Len() int {
	return len(p.pids)
}

// Clear removes all the elements in the set.
func (p *PIDSet) Clear() {
	p.pids = p.pids[:0]
	p.lookup = make(map[int64]IPid)
}

// Empty reports whether the set is empty.
func (p *PIDSet) Empty() bool {
	return p.Len() == 0
}

// Values returns all the elements of the set as a slice.
func (p *PIDSet) Values() []IPid {
	return p.pids
}

// ForEach invokes f for every element of the set.
func (p *PIDSet) ForEach(f func(i int, pid IPid)) {
	for i, pid := range p.pids {
		f(i, pid)
	}
}

func (p *PIDSet) Get(index int) IPid {
	return p.pids[index]
}

func (p *PIDSet) Clone() *PIDSet {
	return NewPIDSet(p.pids...)
}
