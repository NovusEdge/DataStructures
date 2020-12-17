#include "iostream"
#include "vector"

class Node;
class LinkedList;

class LinkedList{
    public:
        Node* head;

    LinkedList(Node* head) { this->head = head; }
    ~LinkedList();

    void Append(int);   
    void Insert(int, int);
    int Pop();
    int Delete(int);

};

class Node{
    public:
        int data;
        Node* next = NULL;
    
    Node(int data, Node* next);
    void PrintNode();
};

void LinkedList::Append(int elem){
    if (this->head->next == NULL){
        this->head->next = new Node(elem, NULL);    
    }

    Node* tempNode = this->head;

    while (tempNode->next != NULL) {
        tempNode = tempNode->next;
    }

    tempNode->next = new Node(elem, NULL);
}

LinkedList MakeList(std::vector<int> items){
    Node* headNode = new Node(items[0], NULL);
    LinkedList retList = LinkedList(headNode);

    for(std::vector<int>::size_type i = 1; i != items.size(); i++) {
        retList.Append(items[i]);    
    }

    return retList;
}

int LinkedList::Pop(){
    
    return 0;
}

int LinkedList::Delete(int index){

    return 0;
}

void Node::PrintNode(){

}

int main() {
    return 0;
}