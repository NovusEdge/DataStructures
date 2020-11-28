class Stack:
    
    def __init__(self):
        self.items = [ ]
    
    def __len__(self):
        return len( self.items )
    
    def __getitem__(self, index):
        return self.items.__getitem__(index)
    
    def __repr__(self):
        return self.items
    
    def __eq__(self, other):
        return self.items == other.items if isinstance( other, Stack ) else None
    
    def pop(self, index=-1):
        return self.items.pop( index )
    
    def push(self, item):
        self.items.append( item )

    def flush(self):
        self.items = [ ]
