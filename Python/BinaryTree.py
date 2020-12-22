from functools import *


class Node:

    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
        self.parent = None

    def __repr__(self):
        return hex(id(self))

    def __str__(self):
        return f"( { self.__repr__() } : { self.data } )"


def BST(object):

    def __init__(self, root=Node()):
        self.root = root

    def insert(self, node):
        self._insert(node, self.root)


    def _printTree(self):
        self.printTree()

    def _insert(self, node, cur_node):

        if node.data < cur_node.data:
            if cur_node.data.left == None:
                cur_node.left = node
            else:
                self._insert(node, cur_node.left)

        else:
            if cur_node.data.right == None:
                cur_node.right = node
            else:
                self._insert(node, cur_node.right)

    def printTree(self):
        pass

if __name__ == "__main__":
    b = BST(Node(15))
    b.insert(Node(10))
    b.insert(Node(20))
    print( b.root.data )
    print( b.root.left.data )
    print( b.root.right.data )
