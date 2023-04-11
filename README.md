# Data Serialization Experiment


```zsh
$ go run cmd/full-experiment/main.go
```
```
Size
+--------+----------+-------------+-------------------+
| #      | JSON     | MESSAGEPACK |  PROTOCOL BUFFERS |
+--------+----------+-------------+-------------------+
| Tiny   | 2.0KiB   | 1.6KiB      | 1.0KiB            |
| Small  | 30.4KiB  | 23.8KiB     | 15.3KiB           |
| Medium | 306.5KiB | 241.6KiB    | 154.0KiB          |
| Large  | 10.1MiB  | 8.0MiB      | 5.1MiB            |
| Huge   | 514.4MiB | 403.8MiB    | 258.8MiB          |
+--------+----------+-------------+-------------------+
serializing time
+--------+--------------+--------------+-------------------+
| #      |         JSON |  MESSAGEPACK |  PROTOCOL BUFFERS |
+--------+--------------+--------------+-------------------+
| Tiny   |    140.417µs |     21.542µs |         274.166µs |
| Small  |    113.541µs |    104.292µs |          68.333µs |
| Medium |      1.358ms |    926.708µs |         958.333µs |
| Large  |  22.006833ms |   17.09325ms |       10.971291ms |
| Huge   | 1.224952458s | 1.035652209s |      634.592666ms |
+--------+--------------+--------------+-------------------+
deserializing time
+--------+--------------+-------------+-------------------+
| #      |         JSON | MESSAGEPACK |  PROTOCOL BUFFERS |
+--------+--------------+-------------+-------------------+
| Tiny   |     79.166µs |    51.833µs |           6.709µs |
| Small  |    516.833µs |   153.708µs |          79.084µs |
| Medium |   5.587542ms |  1.516833ms |         728.958µs |
| Large  |  93.901542ms | 29.175667ms |        14.91575ms |
| Huge   | 4.823344208s | 1.57175175s |      791.295041ms |
+--------+--------------+-------------+-------------------+
```
