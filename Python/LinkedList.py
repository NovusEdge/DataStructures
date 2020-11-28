class Node:
    
    def __init__(self, value=None):
        self.value = value
        self.next = None
        
    def __repr__(self):
        return hex(id(self))
    
    def __str__(self):
        return f"({ self.__repr__() } : { self.value })"
    





class LinkedList:
    
    def __init__(self, head):
        self.head = head
        
    def __repr__(self):
        start = self.head
        nodeList = []
        
        while start != None:
            nodeList.append( ( repr(start), start.value) )
            start = start.next
        
        formStr = r" {} --> " * len(nodeList)
        return formStr.format(*nodeList) + " (None, None) "

    def __len__(self):
        nodeCounter = 0
        start = self.head
        
        while start != None:
            nodeCounter += 1
            start = start.next

        return nodeCounter
    
    def __iter__(self):
        node = self.head
        while node != None:
            yield node
            node = node.next
         
    
    def tail(self):
        tempNode = self.head
        tailNode = None
        
        while tempNode != None:
            tailNode = tempNode
            tempNode = tempNode.next
            
        return tailNode
        
    
    def append(self, newNode):
        self.tail().next = newNode
        
    def find(self, val):
        tempNode = self.head
        
        while tempNode.data != val:
            tempNode = tempNode.next

        return tempNode
    
    def insert(self, insert_at, nodeObj):
        nodeObj.next = insert_at.next
        insert_at.next = nodeObj
        
    def delete(self, delete_after):
        delete_after.next = delete_after.next.next
        
        
        
        
            
class CircularLinkedList:
    
    def __init__(self, head):
        self.head = head
        self.head.next = self.head
    
    def __repr__(self):
        start = self.head.next
        nodeList = [ self.head ]
        
        while start != self.head:
            nodeList.append( ( repr(start), start.value) )
            start = start.next
        
        formStr = r" {} --> " * len(nodeList)
        return formStr.format(*nodeList) + str(self.head)
    
    def __len__(self):
        nodeCounter = 0
        start = self.head.next
        
        while start != self.head:
            nodeCounter += 1
            start = start.next

        return nodeCounter
    
    def __iter__(self):
        tempNode = self.head
        
        while tempNode.next != self.head:
            yield tempNode
            tempNode = tempNode.next
            
        yield self.head
        
    def insert( self, nodeObj, append_after=None ):
        tempNode = append_after
        if tempNode is None:
            tempNode = self.head
        
        nodeObj.next = tempNode.next
        tempNode.next = nodeObj
        
    
    def delete( self, at_front=True ):
        
        if self.head.next is self.head:
                raise Exception("No Node to delete before the head node")
        
        if at_front:
            self.head.next = self.head.next.next
            
        else:
            refNode = self.head.next
            prevNode = self.head
            
            while refNode.next != self.head:
                refNode, prevNode = refNode.next, prevNode.next
                
                if refNode.next == self.head:
                    prevNode.next = self.head
                    
            else:
                prevNode.next = self.head
                
    def find( self, data ):
        refNode = self.head.next
        while refNode.value != data:
            refNode = refNode.next
        else:
            return refNode
    