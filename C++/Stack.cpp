#include "iostream" 
#include "vector"

template <class T> class Stack{
    private:
        std::vector<T> items;
    
    public:
        Stack()  { }   // constructor
        ~Stack() { }   // destructor

        void Push(T);
        T  Pop();
        T  Peek();
        long  Size();
        bool IsEmpty();
};

template<class T>
void Stack<T>::Push( T elem ){
    this->items.push_back( elem );
}

template<class T>
Stack<T> MakeStack(std::vector<T> arr){
    Stack<T> retStack = Stack<T>();

    for(int i: arr){
        retStack.Push(i);
    }

    return retStack;
}

template<class T>
T Stack<T>::Pop(){
    T lastelem = this->items.back();
    this->items.pop_back();

    return lastelem;
}

template<class T>
T Stack<T>::Peek(){
   return this->items.back();
}

template<class T>
long Stack<T>::Size(){
    return long(this->items.size());
}

template<class T>
bool Stack<T>::IsEmpty(){
    return this->items.size() == 0;
}

int main(){
    Stack<int> st = Stack<int>();

    printf("%d\n", st.IsEmpty());
    st.~Stack();
    
    int arr[] = { 10, 20, 30, 40, 50, 60 }; 
    std::vector<int> vect(arr, arr + 6); 

    Stack<int> s = MakeStack(vect);

    printf("Top: %d\n", s.Peek());
    printf("Size: %ld\n", s.Size());

    printf("Popped elem: %d\n", s.Pop());

    printf("Size: %ld\n", s.Size());
    printf("Top: %d\n", s.Peek());

    return 0;
}