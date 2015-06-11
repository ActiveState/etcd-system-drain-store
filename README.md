# etcd_system_drain_store 
 
It allows to submit a new etcd node with a system app_id under `/loggregator/services/` directory                      

# Getting Started

This tool assumes you are working in a standard Go workspace,
as described in http://golang.org/doc/code.html.

make sure you have godep installed 

`$ go get github.com/tools/godep`

Restore all the dependencies

`$ godep restore`

and then build

`$ go build`

Example usage:     

*etcdUrl is optional  

# To Create a key:

`./etcd_system_drain_store -syslogUrl=syslog://192.168.6.144:9008 -etcdUrl=http://127.0.0.1:4001`                                   
 
output :

```
{
    "action": "get", 
    "node": {
        "createdIndex": 195930, 
        "dir": true, 
        "key": "/loggregator/services/system", 
        "modifiedIndex": 195930, 
        "nodes": [
            {
                "createdIndex": 195930, 
                "key": "/loggregator/services/system/4277ff8cad3df053c47158916910e205e88750d2", 
                "modifiedIndex": 195930, 
                "value": "syslog://192.168.6.144:9003"
            }, 
            {
                "createdIndex": 200401, 
                "key": "/loggregator/services/system/b66e42a8997c4b657b8d502aabfbb41a90459eee", 
                "modifiedIndex": 200401, 
                "value": "syslog://192.168.6.144:9008"
            }
        ]
    }
}

```

# To Delete a key:

`./etcd_system_drain_store -syslogUrl=syslog://192.168.6.144:9008 -delete=true`

output:

```
{
    "action": "get", 
    "node": {
        "createdIndex": 195930, 
        "dir": true, 
        "key": "/loggregator/services/system", 
        "modifiedIndex": 195930, 
        "nodes": [
            {
                "createdIndex": 195930, 
                "key": "/loggregator/services/system/4277ff8cad3df053c47158916910e205e88750d2", 
                "modifiedIndex": 195930, 
                "value": "syslog://192.168.6.144:9003"
            }
        ]
    }
}

```
