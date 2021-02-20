#include <stdio.h>
#include "LinkedList.h"


LList make_LList();
long len(LList* llist);
void f_append(LList* llist, Node node);
void b_append(LList* llist, Node node);


typedef struct Node {
    Node* next;
    Node* prev;
    int data;
}Node;


typedef struct LinkedList {
    Node* head;
    long len;
}LList;


LList make_LList(){
    LList res;
    Node headNode;

    headNode.next = NULL;
    headNode.prev = NULL;

    res.len = 0;
    res.head = &headNode;

    return res;
}

long len(LList* llist){
    return llist->len;
}

void f_append(LList* llist, Node node){
    if (llist->head == NULL) {
        llist->head = &node;
        goto end;
    }

    Node tempNode = *((llist->head)->next);
    Node prevNode = *(llist->head);

    while( (&tempNode) != NULL ) {
        tempNode = (*(tempNode.next)) ;
        prevNode = (*(prevNode.next)) ;
    }

    tempNode.next = &node;
    tempNode.prev = &prevNode;
    goto end;

    end:
}

void b_append(LList* llist, Node node) {
    if (llist->head == NULL){
        (*llist->head) = node;

        goto end;
    }

    (*llist->head).prev = &node;
    (*llist->head).prev->next = llist->head;
    end:
}
