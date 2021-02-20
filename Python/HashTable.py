# class Item:
#     def __init__(self, data):
#         self.addr = hex(id(self))
#         self.data = data
    
class HashTable:
    
    def __init__(self, hashfunc):
        self.hashfunc = hashfunc
        self.items = []
    
    
        