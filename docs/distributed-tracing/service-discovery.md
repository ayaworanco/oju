# Service Discovery

This feature is a main feature when we need to connect one service tracing into another without passing trace-context headers.

## Diagram
![service discovery](/service_discovery.png "Service Discovery diagram")

## Sending trace to Oju server
-   First a service `Worker A` sends a trace to Oju server and inside this trace has all attributes and the service that `Worker A` is calling. Here it is an example of the trace body.
    ```json
    {
        "app_key": "4DCD517E-0B5D-4BF9-B101-E353E3934F28",
        "name": "call-worker-b",
        "service": "0C7E8650-EE65-4859-B8A9-3900625437A3",
        "attributes" : {
            "http.method": "POST",
            "http.url": "http://worker-b.api.svc.cluster.local"
        },
        "children": {}
    }
    ```
-   This will be processed and stored in the stack, and as the first trace is alone, nothing will happen.
    ```
    Stack
    --------------------
    <<- Trace A (alone)
    --------------------
    ```
-   Then another trace will be incoming to Oju server from `Worker B` telling that is the request from `Worker A`. And will be processed and stored in the stack.
    ```json
    {
        "app_key": "0C7E8650-EE65-4859-B8A9-3900625437A3",
        "name": "http-post-worker-b",
        "service": "",
        "attributes" : {
            "http.method": "POST",
            "http.url": "http://worker-b.api.svc.cluster.local"
            "http.attributes.name": "John"
        },
        "children": {}
    }
    ```

## Processing the stack
-   And now the traces will be processed in the stack 
    ```
    Stack
    -----------------------
    <<- Trace A <<- Trace B
    -----------------------

    Trace be is child from Trace A?
    -----------------------
    <<- Trace A <<- Trace B
    -----------------------

    Yes! Now append Trace B in Trace A child list
    -----------------------
    <<- Trace A [children: { Trace B }] <<- Trace B
    -----------------------

    Now remove Trace A from the stack and keep Trace B for later comparison
    -----------------------
    *poof* <<- Trace B
    -----------------------
    ```
