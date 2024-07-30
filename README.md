# A simple demo for PAM plugin using go-plugin

To run the demo:

- go build -o main main.go
- go build -o pam_plugin_main pam_plugin_main.go
- ./main

Sample output:
```
$ ./main                                        
2024-07-30T14:27:29.691-0500 [DEBUG] plugin: starting plugin: path=./pam_plugin_main args=[./pam_plugin_main]
2024-07-30T14:27:29.692-0500 [DEBUG] plugin: plugin started: path=./pam_plugin_main pid=43151
2024-07-30T14:27:29.692-0500 [DEBUG] plugin: waiting for RPC address: plugin=./pam_plugin_main
2024-07-30T14:27:30.513-0500 [DEBUG] plugin: using plugin: version=1
2024-07-30T14:27:30.513-0500 [DEBUG] plugin.pam_plugin_main: plugin address: network=unix address=/var/folders/gm/2s1k7wb1747d5krds0jjcxh40000gn/T/plugin2787231300 timestamp=2024-07-30T14:27:30.513-0500
[
  {
    "name": "user1",
    "password": "password1",
    "createdAt": "2024-07-30T14:27:30.51273-05:00"
  },
  {
    "name": "user2",
    "password": "password2",
    "createdAt": "2024-07-30T14:27:30.51273-05:00"
  }
]
Password rotated for user1.
[
  {
    "name": "user1",
    "password": "NewPassword1",
    "createdAt": "2024-07-30T14:27:30.51273-05:00"
  },
  {
    "name": "user2",
    "password": "password2",
    "createdAt": "2024-07-30T14:27:30.51273-05:00"
  }
]
2024-07-30T14:27:30.515-0500 [DEBUG] plugin.pam_plugin_main: 2024/07/30 14:27:30 [DEBUG] plugin: plugin server: accept unix /var/folders/gm/2s1k7wb1747d5krds0jjcxh40000gn/T/plugin2787231300: use of closed network connection
2024-07-30T14:27:30.515-0500 [INFO]  plugin: plugin process exited: plugin=./pam_plugin_main id=43151
2024-07-30T14:27:30.515-0500 [DEBUG] plugin: plugin exited
```
