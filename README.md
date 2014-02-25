Hashed Wheel Timer
===

Simple hashed wheel timer. See http://www.cs.columbia.edu/~nahum/w6998/papers/sosp87-timing-wheels.pdf

TODO:

1. Add `Size` to inspect the size the `WheelTimer`.
2. Add `Length` to `Node` inspect the length of the linked list.
3. Allow scheduling ticks outside of the immediate interval. 
4. Support at-least concurrent writes (multi-producer, single consumer).
