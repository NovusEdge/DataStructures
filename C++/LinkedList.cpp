#include "iostream"
#include "vector"

template<class T> class Node;
template<class T> class LinkedList;

template<class T>
class LinkedList{

    public:
        Node<T>* head;
        Node<T>* tail;
        int      len;

    LinkedList(Node<T>* head) { this->head = head; }
    ~LinkedList() {};

    void Append(T);   
    void Insert(T, int);
    T    Pop();
    T    Delete(int);

};

template<class T>
class Node{
    public:
        T        data;
        Node<T>* next = NULL;
        
    Node(T data, Node<T>* next)  {this->data = data; this->next = next; }
    ~Node() {};

};

template<class T>
void LinkedList<T>::Append(T elem){
    if (this->head->next == NULL){
        this->head->next = new Node<T>(elem, NULL);    
    }

    Node<T>* tempNode = this->head;

    while (tempNode->next != NULL) {
        tempNode = tempNode->next;
    }

    tempNode->next = new Node<T>(elem, NULL);
    this->tail = tempNode->next;
    this->len += 1;
}

template<class T>
LinkedList<T> MakeList(std::vector<T> items){
    Node<T>* headNode = new Node<T>(items[0], NULL);
    LinkedList<T> retList = LinkedList<T>(headNode);

    for(int i = 1; i != items.size(); i++) {
        retList.Append(items[i]);    
    }

    return retList;
}

template<class T>
T LinkedList<T>::Pop(){
    Node<T>* tempNode = this->head->next;
    Node<T>* prevNode = this->head;

    while ( tempNode->next != NULL ){
        tempNode = tempNode->next;
        prevNode = prevNode->next;
    }

    this->tail = prevNode;
    this->tail->next = NULL;

    T retData = tempNode->data;

    tempNode->~Node();

    this->len -= 1;

    return retData;
}

template<class T>
T LinkedList<T>::Delete(int index){

    if ( index >= this->len ){
        return -1;
    }

    Node<T>* tempNode = this->head->next;
    Node<T>* prevNode = this->head;
    int i = 0;

    while ( i < index ){
        i++;
        tempNode = tempNode->next;
        prevNode = prevNode->next;
    }

    T retData = tempNode->data;
    prevNode->next = tempNode->next;

    tempNode->~Node();

    return retData;
}


int main() {
    Node<int> headNode = Node<int>(10, NULL);
    LinkedList<int> ll = LinkedList<int>(&headNode);

    printf("Head: %p\n", ll.head); 

    return 0;
}