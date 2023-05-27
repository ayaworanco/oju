# Using Qoju

Qoju is the binary that you want to improve your observability on-demand. This will serve as a command query to filtering the logs from what you want

> ## Query example

```sh
qoju '$ipv4 eq 54.36.149.41 and $status_code eq 400'
```

Let's break down this example

1. `$ipv4`, `$status_code` is the super variable that Oju already have mapped with some regular expressions, they only need to match the value in the right-side
2. `eq`, `and` are logical operators, that already pre-mapped, will be used to check the values in left-side to right-side.
