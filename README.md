# ls-json
List information about the FILE in json format (the current directory by default)

Usage:
  ls-json [FILE] [flags]

```
» ls-json
{"ext":"","name":"HEAD","path":"/home/rok/src/github.com/aca/ls-json/.git/HEAD","bytes":23,"size":"23 B","mod_time":"2020-06-27T21:04:39.995260916+09:00"}
{"ext":"","name":"config","path":"/home/rok/src/github.com/aca/ls-json/.git/config","bytes":285,"size":"285 B","mod_time":"2020-06-27T21:04:39.995260916+09:00"}
{"ext":"","name":"description","path":"/home/rok/src/github.com/aca/ls-json/.git/description","bytes":73,"size":"73 B","mod_time":"2020-06-27T21:04:37.631927418+09:00"}
{"ext":".sample","name":"applypatch-msg.sample","path":"/home/rok/src/github.com/aca/ls-json/.git/hooks/applypatch-msg.sample","bytes":478,"size":"478 B","mod_time":"2020-06-27T21:04:37.631927418+09:00"}
{"ext":".sample","name":"commit-msg.sample","path":"/home/rok/src/github.com/aca/ls-json/.git/hooks/commit-msg.sample","bytes":896,"size":"896 B","mod_time":"2020-06-27T21:04:37.631927418+09:00"}
{"ext":".sample","name":"fsmonitor-watchman.sample","path":"/home/rok/src/github.com/aca/ls-json/.git/hooks/fsmonitor-watchman.sample","bytes":4655,"size":"4.7 kB","mod_time":"2020-06-27T21:04:37.631927418+09:00"}
...
```

```
» ls-json ls-json.go go.sum
{"ext":".go","name":"ls-json.go","path":"/home/rok/src/github.com/aca/ls-json/ls-json.go","bytes":1953,"size":"2.0 kB","mod_time":"2020-06-27T21:27:55.612024472+09:00"}
{"ext":".sum","name":"go.sum","path":"/home/rok/src/github.com/aca/ls-json/go.sum","bytes":12575,"size":"13 kB","mod_time":"2020-06-27T20:53:36.071881489+09:00"}
```

```
» ls-json | jq 'select(.ext? | match("go"))'
{
  "ext": ".go",
  "name": "ls-json.go",
  "path": "/home/rok/src/github.com/aca/ls-json/ls-json.go",
  "bytes": 1953,
  "size": "2.0 kB",
  "mod_time": "2020-06-27T21:27:55.612024472+09:00"
}
```
