package main

import "log"

func main() {

	stackQueue := NewQueue()
	stackQueue.enQueue(1)
	stackQueue.enQueue(2)
	stackQueue.enQueue(3)
	log.Println("deQueue value is",stackQueue.deQueue())
	log.Println("deQueue value is",stackQueue.deQueue())
	stackQueue.enQueue(4)
	log.Println("myQueue head value is",stackQueue.peek())
	log.Println("deQueue value is",stackQueue.deQueue())
	log.Println("myQueue is empty:",stackQueue.isEmpty())
	log.Println("deQueue value is",stackQueue.deQueue())
	log.Println("myQueue is empty:",stackQueue.isEmpty())
}
