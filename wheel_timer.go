package wheel_timer

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head *Node
}

type WheelTimer struct {
	maxInterval int
	currentTick int

	ring []*List
}

func New(maxInterval int) *WheelTimer {
	ring := make([]*List, maxInterval)
	for i := range ring {
		ring[i] = new(List)
	}
	return &WheelTimer{
		maxInterval: maxInterval,
		ring:        ring,
	}
}

func (w *WheelTimer) Schedule(ticks int, value interface{}) {
	node := &Node{Value: value}
	index := (w.currentTick + ticks) % w.maxInterval
	node.Next = w.ring[index].Head
	w.ring[index].Head = node
}

func (w *WheelTimer) Tick() (node *Node) {
	node = w.ring[w.currentTick].Head
	w.ring[w.currentTick].Head = nil
	w.currentTick = (w.currentTick + 1) % w.maxInterval
	return
}
