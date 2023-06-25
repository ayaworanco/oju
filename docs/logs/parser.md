# Log parser

This log parser is based on a paper called [Drain](https://jiemingzhu.github.io/pub/pjhe_icws2017.pdf) <br>
It's based on a fixed depth tree that consists in break the log message into nodes and create a pattern giving the log a template.

## Drain Pipeline

![pipeline](/drain_pipeline.png "Pipeline diagram")

-  The log message will pass through a tokenization and set what type of length this log is.
-  The parser will create nodes inside the layer checking if that token has numbers, if has numbers this will be a wildcard
-  Then we will have a Log Group in the leaf node that will hold the Log Event (that will be the template) and the parameters

## Different logs
![pipeline](/drain_new_logs.png "Drain new Logs")

-  When a new log cames with the same Length, it just start a new branch with the same idea
-  That will give a way to make a standard log events even if these logs are unstructured