# Running

First of all you need a config file called `config.json` in your working directory within `oju` binary.

## Config file example

An example of the `config.json` file

```json
{
    "allowed_applications": [
       {
            "name": "worker-1",
            "app_key": "<app-generated-key>", // Can be an UUID
            "host": "http://worker.api.svc.cluster.local" // Example if is stored on kubernetes
       } 
    ]
}
```

## Using the binary

In your working directory with your `config.json` you just run
```sh
./oju
```

If you are in MacOS you probably will need to allow permission for the binary run.

<!--

The AWO protocol will be the communication that we will use within Oju. He is quitely similar to HTTP protocol with a few differences. AWO is used only for this purpose, entry logs and queries.

An example of a message to entry log
```
LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO1.1
54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/2713%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""
```

This structure can be separeted by
- Header -> LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO1.1
- Timer -> 02:49:12
- Log Message

Some other things are important here, like the UUID key here, is called `app-key` that will help to find and query entries only for this application, and any message to Oju with a unrecognized app-key will be rejected.
So to understande the structure

```
<AWO VERB> <APP-KEY> <AWO VERSION>
<MESSAGE>
```

## Example using a lang

Ruby:
```ruby
require 'socket'

socket = TCPSocket.new('localhost', 8080)
message = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO1.1\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/2713%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

socket.write(message.strip)
socket.close
```

This will send a message and all things will be processed
-->
