#include "iostream" 
#include "vector"

class Stack{
    private:
        std::vector<int> items;
    
    public:
        Stack()  {}   // constructor
        ~Stack() {}   // destructor

        void Push(int);
        int  Pop();
        int  Peek();
        int  Size();
        bool IsEmpty();
};

void Stack::Push( int elem ){
    this->items.push_back( elem );
}

Stack MakeStack(std::vector<int> arr){
    Stack retStack = Stack();

    for(int i: arr){
        retStack.Push(i);
    }

    return retStack;
}

int Stack::Pop(){
    int lastelem = this->items.back();
    this->items.pop_back();

    return lastelem;
}

int Stack::Peek(){
   return this->items.back();
}

int Stack::Size(){
    return this->items.size();
}

bool Stack::IsEmpty(){
    return Size() == 0;
}

int main(){
    Stack st = Stack();

    printf("%d\n", st.IsEmpty());
    st.~Stack();
    
    int arr[] = { 10, 20, 30, 40, 50, 60 }; 
    std::vector<int> vect(arr, arr + 6); 

    Stack s = MakeStack(vect);

    printf("Top: %d\n", s.Peek());
    printf("Size: %d\n", s.Size());

    printf("Popped elem: %d\n", s.Pop());

    printf("Size: %d\n", s.Size());
    printf("Top: %d\n", s.Peek());

    return 0;
}