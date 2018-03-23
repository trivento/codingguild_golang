# Gossip

Let's create our own gossip network.
We need one node in the network initially.
Then we can start others node, which will make it self known to one of the existing nodes.

All nodes will keep sending each other the list of known hosts periodically.

```
mkdir $HOME/go
export GOPATH=$HOME/go
go get github.com/trivento/codingguild_golang/gossip
```

Start the main node

    $GOPATH/bin/gossip -port 9090


Start a second node, which makes itself known to the main node.

    $GOPATH/bin/gossip -port 9091 -seednode http://HOST:PORT

Start some more nodes. 
In the current implementation, this could break the main node! (fatal error: concurrent map writes)


    SEEDNODE=http://HOST:PORT
    port=9100
    while [ $port -lt 9105 ]
    do
        $GOPATH/bin/gossip -port $port -seednode $SEEDNODE &
        port=$[$port+1]
    done

The list of hosts will now propagate to the netwokr.



