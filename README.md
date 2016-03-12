# mmic
MatterMost Integration Center.

MMIC support multiple service (eg. trello, github) webhook integration.

##
```
+--------------+
|   github     +--------+
+--------------+        |          +---------------------+           +------------------------+
                        |          |       MMIC          |           |         MatterMost     |
+--------------+        +--------> |/github              |     +---> |/gihubChannel           |
|    trello    +--------+          |                     |     |     |                        |
+--------------+        +--------> |/trello              |     +---> |/trelloStatusChannel    |
                                   |                     +-----+     |                        |
+--------------+        +--------> |/gitlab              |           |                        |
|    gitlab    +--------+          |                     |           |                        |
+--------------+                   |                     |           |                        |
                                   |/.....               |           |                        |
+--------------+                   |                     |           |                        |
|   ......     |                   +---------------------+           +------------------------+
+--------------+
```
