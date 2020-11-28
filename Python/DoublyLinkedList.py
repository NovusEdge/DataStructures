class Node:
    
    def __init__(self, value=None):
        self.value = value
        self.next = None
        self.prev = None
        
    def __repr__(self):
        return hex(id(self))
    
    def __str__(self):
        return f"({ self.__repr__() } : { self.value })"
    
    
    
    
    
    
class DLinkedList:
    
    def __init__(self, head=None):
        self.head = head
    
    def __repr__(self):
        start = self.head
        nodeList = []
        
        while start != None:
            nodeList.append( ( repr(start), start.value) )
            start = start.next
        
        formStr = r" {} <-> " * len(nodeList)
        return formStr.format(*nodeList) + " (None, None) "
        
    def __iter__(self):
        node = self.head
        while node != None:
            yield node
            node = node.next

    def __len__(self):
        nodeCounter = 0
        start = self.head
        
        while start != None:
            nodeCounter += 1
            start = start.next
            
        return nodeCounter

    def tail(self):
        tempNode = self.head
        tailNode = None
        
        while tempNode != None:
            tailNode = tempNode
            tempNode = tempNode.next
            
        return tailNode
    
    def f_append(self, nodeObj):
        nodeObj.prev = self.tail()
        self.tail().next = nodeObj
             
    def b_append(self, nodeObj):
        self.head.prev = nodeObj
        nodeObj.next = self.head
        self.head = nodeObj
        
    def delete(self, data):
        tempNode = self.head
        
        while tempNode != None:
            
            if tempNode.value == data:
                tempNode.prev.next = tempNode.next
                tempNode.next = tempNode.prev
            
            tempNode = tempNode.next
        
        
        
        
        
        
class DCircularLinkedList:
    
    def __init__(self, head):
        self.head = head
        self.head.prev = self.head
        self.head.next = self.head
    

    def __repr__(self):
        start = self.head.next
        nodeList = [ self.head ]
        
        while start != self.head:
            nodeList.append( ( repr(start), start.value) )
            start = start.next
        
        formStr = r" {} <-> " * len(nodeList)
        return formStr.format(*nodeList) + str(self.head)

    def __iter__(self):
        tempNode = self.head
        
        while tempNode.next != self.head:
            yield tempNode
            tempNode = tempNode.next
            
        yield self.head
    
    def __len__(self):
        nodeCounter = 0
        start = self.head.next
        
        while start != self.head:
            nodeCounter += 1
            start = start.next

        return nodeCounter
            
    def append(self, nodeObj, at_front=True):
        if at_front:
            nodeObj.next = self.head.next
            nodeObj.prev = self.head
            nodeObj.next.prev = nodeObj
            self.head.next = nodeObj
        else:
            nodeObj.prev = self.head.prev
            self.head.prev.next = nodeObj
            self.head.prev = nodeObj
            nodeObj.next = self.head
        
    
    def delete(self, at_front=True):
        pass
        
