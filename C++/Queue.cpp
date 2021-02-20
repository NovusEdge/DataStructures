#include "iostream"
#include "vector"

template<class T> class Queue{
    private:
        std::vector<T> items;
	
	public:

	    Queue(std::vector<T> items)    { this->items = items; };
    	~Queue()   {};

    	long Len( );
	    void PrintQueue( );
		void Enqueue(T elem);
		T Dequeue();
};


template<typename T>
void Queue<T>::PrintQueue(){
	printf(" <-- ");

	for(long i = 0; i < this->items.size(); i++){
		std::cout<< " | " << this->items[i] << " | ";
	}


	printf(" --> \n");
}


template<class T>
long Queue<T>::Len() {
	return long(this->items.size());
}

template<class T>
void Queue<T>::Enqueue(T elem){
	this->items.push_back(elem);
}

template<class T>
T Queue<T>::Dequeue(){
	T temp = *(this->items.begin());
	this->items.erase(this->items.begin());
	return temp;
}

int main() {
    int arr[] = { 10, 20, 30, 40, 50, 60 }; 
    std::vector<int> vect(arr, arr + 6); 
	Queue<int> q = Queue<int>(vect);
	
	q.Enqueue(10);
	q.PrintQueue();

	printf("%d\n", q.Dequeue());
	q.PrintQueue();
	
	printf("Length: %ld\n", q.Len());

	return 0;
}
