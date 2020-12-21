#include "iostream"
#include "vector"


template<class T>
class Node{

	public:
		T	  data;
		Node* left  = NULL;
		Node* right = NULL;
}



template<class T>
class BinaryTree{
	public:
		Node<T>* root = NULL;

		BianryTree() {};
		~BinaryTree() {};

		void AddNode(Node<T>* node, Node<T>* curNode);
		Node<T>* Search(T data);
		Node<T>* Delete(T data);

}

template<class T>
Node<T>* BinaryTree<T>::Search(T data){
	Node<T>* tempNode = this->root;
	
	while (1){
		
		if (data == tempNode->data || (tempNode->left == NULL && tempNode->right == NULL)){
			return tempNode
		}

		if ( tempNode->left->data < data ){
			if (tempNode->left == NULL){
				tempNode = tempNode->right;
			}
			else if(tempNode->right == NULL){
			tempNode = tempNode->left;
			}
		}
		else{
			if (tempNode->right == NULL){
				tempNode = tempNode->left;
			} 
			else if (tempNode->right == NULL){
				tempNode = tempNode->right;
			}
		}
	}
}

template<class T>
void AddNode(Node<T>* node, Node<T>* curNode ){

}




