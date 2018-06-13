"""
MEDIUM

Implement a load balancer for web servers. It provide the following functionality:

1. Add a new server to the cluster => add(server_id).
2. Remove a bad server from the cluster => remove(server_id).
3. Pick a server in the cluster randomly with equal probability => pick().

Example:
At beginning, the cluster is empty => {}.

add(1)
add(2)
add(3)
pick()
>> 1         // the return value is random, it can be either 1, 2, or 3.
pick()
>> 2
pick()
>> 1
pick()
>> 3
remove(1)
pick()
>> 2
pick()
>> 3
pick()
>> 3

Ideal:
use a list to store server ID, and a map(id=>index in list) to get access when want to remove.
"""
class LoadBalancer:
    def __init__(self):
        # do intialization if necessary
        self.serversIdIndex = {}
        self.serversId = []

    """
    @param: server_id: add a new server to the cluster
    @return: nothing
    """
    def add(self, server_id):
        # write your code here
        if server_id in self.serversIdIndex:
            return
        self.serversId.append(server_id)
        self.serversIdIndex[server_id] = len(self.serversId) - 1
        return

    """
    @param: server_id: server_id remove a bad server from the cluster
    @return: nothing
    """
    def remove(self, server_id):
        # write your code here
        if server_id not in self.serversIdIndex:
            return
        removeServerIndex = self.serversIdIndex[server_id]
        del self.serversIdIndex[server_id]
        lastServerId = self.serversId[-1]
        self.serversIdIndex[lastServerId] = removeServerIndex
        self.serversId[removeServerIndex] = lastServerId
        self.serversId.pop()
        return

    """
    @return: pick a server in the cluster randomly with equal probability
    """
    def pick(self):
        # write your code here
        from random import randint
        index = randint(0, len(self.serversId) - 1)
        return self.serversId[index]
