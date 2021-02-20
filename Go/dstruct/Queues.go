package dstructs

type qItem struct {
	Data interface{}
}

// type dqItem struct {
// 	Data interface{}
// }

//Queue : Creates a queue
type Queue struct {
	Items []qItem
}

// type DQueue struct {
// 	Items []dqItem
// }

//Len : Reports the size of a queue
func (q *Queue) Len() int {
	return len(q.Items)
}

//Enqueue : Performs an enqueue operation on the queue
func (q *Queue) Enqueue(data interface{}) {
	eqItem := qItem{Data: data}
	tempItems := []qItem{eqItem}
	tempItems = append(tempItems, q.Items...)
	q.Items = tempItems
}

//Dequeue : Performs an dequeue operation on the queue
func (q *Queue) Dequeue(data interface{}) {
	if !q.IsEmpty() {
		q.Items = q.Items[:len(q.Items)-1]
	}
}

//IsEmpty : Returns true if the said queue has no items
func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}
