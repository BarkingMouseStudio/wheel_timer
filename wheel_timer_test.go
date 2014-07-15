package wheel_timer

import (
	"fmt"
	"testing"
)

func TestWheelTimer_New(t *testing.T) {
	timer := New(1)
	if timer.ring[0] == nil {
		t.Fatal("Ring not initialized")
	}
}

func TestWheelTimer_Schedule(t *testing.T) {
	timer := New(10)
	timer.Schedule(3, "Tick")

	timer.Tick()
	timer.Tick()
	timer.Tick()
	node := timer.Tick()

	if node == nil {
		t.Fatal("Node undefined")
	}

	if node.Value.(string) != "Tick" {
		t.Fatal("Unexpected value")
	}
}

func TestWheelTimer_Tick(t *testing.T) {
	timer := New(10)
	for i := 0; i < 10; i++ {
		timer.Schedule(i, i)
	}

	for i := 0; i < 10; i++ {
		node := timer.Tick()
		if node == nil {
			t.Fatal("Node undefined")
		}

		if node.Value.(int) != i {
			t.Fatal("Unexpected value", node.Value, i)
		}
	}
}

func BenchmarkWheelTimer_drain(b *testing.B) {
	maxInterval := 20
	timer := New(maxInterval)

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			timer.Schedule(j%maxInterval, i+(j%maxInterval))
		}

		for node := timer.Tick(); node != nil; node = node.Next {
		}
	}
}

func BenchmarkWheelTimer_fill(b *testing.B) {
	maxInterval := 20
	timer := New(maxInterval)

	for i := 0; i < b.N; i++ {
		timer.Schedule(i%maxInterval, (i % maxInterval))
	}
}

func BenchmarkWheelTimer_fast(b *testing.B) {
	maxInterval := 20
	timer := New(maxInterval)

	for i := 0; i < b.N; i++ {
		timer.Schedule(i%maxInterval, (i % maxInterval))
		timer.Tick()
	}
}

func ExampleWheelTimer() {
	timer := New(10)

	for tick := 0; tick < 10; tick++ {

		for interval := 0; interval < 2; interval++ {
			timer.Schedule(interval, fmt.Sprintf("Tick: %v", tick+interval))
		}

		for node := timer.Tick(); node != nil; node = node.Next {
			fmt.Println(tick, node.Value)
		}
	}

	// Output:
	// 0 Tick: 0
	// 1 Tick: 1
	// 1 Tick: 1
	// 2 Tick: 2
	// 2 Tick: 2
	// 3 Tick: 3
	// 3 Tick: 3
	// 4 Tick: 4
	// 4 Tick: 4
	// 5 Tick: 5
	// 5 Tick: 5
	// 6 Tick: 6
	// 6 Tick: 6
	// 7 Tick: 7
	// 7 Tick: 7
	// 8 Tick: 8
	// 8 Tick: 8
	// 9 Tick: 9
	// 9 Tick: 9
}
