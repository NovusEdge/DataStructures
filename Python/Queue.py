class Queue:
    
    def __init__(self):
        self.items = []
        
    def __len__(self):
        return len(self.items)
    
    def __repr__(self):
        return self.items
    
    def __getitem__(self, index):
        return self.items.__getitem__(index)
    
    def __eq__(self, other):
        return self.items == other.items if isinstance(other, Queue) else None
    
    def enqueue(self, item):
        self.items.insert(0, item)
        
    def dequeue(self):
        return self.items.pop()
    
    def clear(self):
        self.items = []
        





class DQueue:
    
    def __init__(self):
        self.items = [ ]
    
    def __repr__(self):
        return self.items
    
    def __str__(self):
        return self.items
    
    def __len__(self):
        return self.items.__len__()
    
    def __getitem__(self, index):
        return self.items[ index ]
    
    def __eq__(self, other):
        return self.items == other.items if isinstance( other, Queue ) else None

    def enqueue(self, item, at_front=True):
        if at_front:
            self.items.append(item)
        else:
            self.items.insert(0, item)
            
    def dequeue(self, at_front=True):
        if at_front:
            return self.items.pop()
        
        else:
            return self.items.pop(0)
        
    def clear(self):
        self.items = [ ]