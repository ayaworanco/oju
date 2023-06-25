# AWO Protocol

This protocol resumes how your client will talk to Oju server. <br>
AWO has a simple TCP message template to follow:

```
<VERB> <APP-KEY> <VERSION>\n
<MESSAGE>
```

## Explaning some objects

-  `VERB` is the action that you want to talk with Oju
  - Here it is some available verbs:
    -  `LOG` is for sending logs that Oju will parse and structure
    -  `TRACE` is for sending distributed traces
-  `APP-KEY` is the application key that you want to generate to keep secure the message exchange between your client and Oju server. This can be an simple text, or an UUID and is better that only your client (like in a env var) and you Oju instance (in your config.json) know each other application keys.
-  `VERSION` is the current version of this protocol and here it is the available versions to put in headers:
    -  `AWO`
    -  `AWO1.1`
-  `MESSAGE` this will be variant by the verb that you used. For example:
    - If you send a `LOG` you have to put in message the output of your log and Oju will parse and structure your logs
    - If you send a `TRACE` you have to follow this [trace template](#) in your messages
