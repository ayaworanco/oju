# Tracer

The tracer object will be represented by this template when stored

```json
{
    "app_key": "<app-key>",
    "name": "<span-name>",
    "service": "<service-key-or-host>",
    "attributes": {
        // some attributes here like
        "http.method": "POST"
    },
    "children": {
        // some another traces
    },
}
```

## Explaning some objects

-   `app-key` is the application key that this trace was generated. This is automatically retrieved by the TCP packet header.
-   `name` is the span name that you want to this trace like: `calculate-something` 
-   `service` is a important one, is the application key or host that you want to be the destination of this trace, like if you send a http request to another service and put into service field the url of the service or the application key.
-   `attributes` is all attributes that you want to fill.
-   `children` is all traces that is child of this one, and will be used when exported to an API too.
