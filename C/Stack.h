#ifndef STACK_H_
#define STACK_H_

typedef struct Stack Stack;

Stack* MakeStack(size_t size);

void Push(Stack* stack, int item);
int Pop(Stack* stack);
void DeleteStack(Stack* stack);

#endif
