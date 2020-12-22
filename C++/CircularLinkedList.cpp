#include "iostream"
#include "vector"

template<class T>
class CLinkedList{
	private:
		std::vector<T> items;
	
	public:
		CLinkedList(std::vector<T> items) { this->items = items }; 
		~CLinkedList(){};
		
		long Len();
		void Enqueue(T item);
		T 	 Dequeue();
	
}
