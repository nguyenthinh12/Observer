package observer

/*Subject is interface*/
type Subject interface{
	Register(Observer observer)
	DeRegister(Observer observer)
	NotifyAll()
}
