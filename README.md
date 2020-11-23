# todo server

user: demo password: 1234

This is the server for the example app. This repo is made for educational purposes only. It is not our idea to show how
to build a server. Our intention is showing the connections of all parts.

The build scripts are calling some other repositorys directly (../todo-client, ../todo-specs). We omitted the usage of a
build pipeline for simplicity.

### Overview

This server exposes the grpc API for the grpc-gateway (transcoder). Usually this would be all a server has to do.

For simplicity the transcoder and the client is also built in.

```

    :10000
   +-------------------------+
   |                         |
   |       todo-client       |  <-----------------+
   |                         |                    |
   +-------------------------+                    |
                |                         +----------------+
    :7001       v                         |   todo-specs   |
   +-------------------------+            +----------------+
   |                         |                    |
   |       grpc gateway      |  <-----------------+
   |                         |                    |
   +-------------------------+                    |
                |                                 |
    :7000       v                                 |
   +-------------------------+                    |
   |                         |                    |
   |       todo-server       |  <-----------------+
   |     ----------------    |
   +-------------------------+



```