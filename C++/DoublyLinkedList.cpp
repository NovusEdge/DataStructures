#include "iostream"
#include "vector"

template<class T> class DNode;
template<class T> class DoublyLinkedList;

template<class T>
class DoublyLinkedList{    // ! DO NOT USE, HAVENT MANAGD MEMORY STUFFS YET

    public:
        DNode<T>* head;
        int len;

    DoublyLinkedList(DNode<T>* head) { this->head = head; }
    ~DoublyLinkedList() { }

    void Append(T);   
    void Insert(T, int);
    T    Pop();
    T    Delete(int);

};

template<class T>
class DNode{
    public:
        T data;
        DNode<T>* next = NULL;
        DNode<T>* prev = NULL;
        
    DNode(T data, DNode<T>* prev, DNode<T>* next)  {this->data = data; this->next = next; this->prev = prev; }
    ~DNode() { }

};


template<class T>
void DoublyLinkedList<T>::Append(T elem){
    if (this->head->next == NULL){
        this->head->next = new DNode<T>(elem, this->head, NULL);
    }

    DNode<T>* tempNode = this->head;

    while (tempNode->next != NULL) {
        tempNode = tempNode->next;
    }

    tempNode->next = new DNode<T>(elem, tempNode, NULL);
    
    this->len += 1;
}

template<class T>
DoublyLinkedList<T> MakeList(std::vector<T> items){
    DNode<T>* headNode = new DNode<T>(items[0], NULL, NULL);
    DoublyLinkedList<T> retList = DoublyLinkedList<T>(headNode);

    for(int i = 1; i != items.size(); i++) {
        retList.Append(items[i]);    
    }

    return retList;
}

template<class T>
T DoublyLinkedList<T>::Pop(){
    DNode<T>* tempNode = this->head;

    while ( tempNode->next != NULL ){
        tempNode = tempNode->next;
    }

    T retData = tempNode->data;

    tempNode->prev->next = NULL;
    tempNode->~Node();

    this->len -= 1;

    return retData;
}

template<class T>
T DoublyLinkedList<T>::Delete(int index){

    if ( index >= this->len ){
        return -1;
    }

    DNode<T>* tempNode = this->head;
    
    int i = 0;

    while ( i < index ){
        tempNode = tempNode->next;
        i++;
    }

    T retData = tempNode->data;

    tempNode->next->prev = tempNode->prev;
    tempNode->prev->next = tempNode->next;

    tempNode->~Node();

    return retData;
}


int main() {
    DNode<int> headNode = DNode<int>(10, NULL, NULL);
    DoublyLinkedList<int> ll = DoublyLinkedList<int>(&headNode);

    printf("Head: %p\n", ll.head); 

    return 0;
}