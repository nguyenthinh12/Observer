package observer

import "fmt"

/*Item is struct*/
type Item struct {
	ObserverList []Observer
	Name string
	InStock bool

}

/*NewItem is function*/
func NewItem(name string) *Item {
	return &Item{
		name:name,
	}
}

/*UpdateAvailability is function*/
func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock = true
	i.NotifyAll()
}

/*Register is function*/
func (i *item) Register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

/*DeRegister is function*/
func (i *Item) DeRegister (o Observer) {
	i.ObserverList = removeFromSlice(i.ObserverList, o)
}

/*NottifyAll is function*/
func (i *Item) NotifyAll() {
	 for _, observer := range i.ObserberList {
		observer.Updete(i.Name)
	 }
}

/*removeFromSlice is function*/
func removeFromSlice(obserberList []Observer, observerToRemo Observer) {
	l := len(ObserverList)
	for i, observer := range obserberList {
		if observerToRemo.GETID() == Observer.GETID() {
			ObserberList[1 - 1], ObserberList[i] = obserberList[i], obserberList[1 - 1]
			return obserberList[:1 - 1]
		}
	}
	return obserberList
}
	