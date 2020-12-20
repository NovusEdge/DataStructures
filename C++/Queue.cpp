#include "iostream"

template<class T>
class Queue{
    private:
        std::vector<T> items;

    Queue(std::vector<T> items)    { this->items = items; };
    ~Queue()   {};

    void Enqueue(T item);
    T    Dequeue();

    long Len( );
    void PrintQueue( );
	// void Enqueue(T elem);

};


template<class T>
void Queue<T>::PrintQueue(){
	printf(" <-- ");
	for (T i: this->items){
		std::cout<< " | " << i << " | ";
	}
	printf(" --> \n");
}


template<class T>
long Queue<T>::Len() {
	return long(this->items.size());;
}


template<class T>


int main() {
	
	return 0;
}
