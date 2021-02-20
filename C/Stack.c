#include <stdlib.h>
#include <stdio.h>
#include "Stack.h"


typedef struct Stack {
    size_t size;
    int* items;
    size_t _count;
} Stack ;

Stack* MakeStack(size_t size);
void Push(Stack* stack, int item);
int Pop(Stack* stack);
void DeleteStack(Stack* stack);

Stack* MakeStack(size_t size){
    Stack* result;
    int* items = (int*)malloc(sizeof(int) * size);

    result->size = size;
    result->items = items;
    result->_count = 0;

    return result;
}

void Push(Stack* stack, int item){

    Stack _stack = *stack;
    size_t _size = _stack.size;

    if (_size == _stack._count) {
        printf("Stack Overflow\n");
        exit(1);
    }
    int* _temp = _stack.items;

    for( int i = 0; i < _stack._count - 1; i++ ){ _temp++; }

    _temp = &item;
    _stack.items = _temp;
    _stack._count++;

    stack = &_stack;
}

int Pop(Stack* stack){

    Stack _stack = *stack;
    size_t _size = _stack.size;
    if (_size == 0){
        printf("Stack Underflow\n");
        exit(1);
    }

    int* _temp = _stack.items;
    for( int i = 0; i < _stack._count - 1; i++ ){ _temp++; }

    int res = *_temp;
    free(_temp);
    _stack._count--;
    stack = &_stack;

    return res;
}

void DeleteStack(Stack* stack){
    free(stack->items);
    free(stack);
    printf("Stack Deleted\n");
}
