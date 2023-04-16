<center> <h1>
  Oluwoye
  </h1>  </center>

<p align="center">
  <img src="images/cowries.png" />
</p>

A stream processor that can be used for log analysis and network security monitoring based on rule system.

> ### Table of contents

- [Contributing](#contributing)
- [Writing Rules](#writing-rules)
- [License](#license)

## Writing Rules

Usually you will need a rules.yaml in root folder that Oluwoye binary is. <br>
An example of creating this rules

```yaml
- resource: $ipv4
  operator: equal
  target: 192.168.1.1
  action:
    name: alert_by_email
    parameters:
      - tech_lead@gmail.com
      - warning
- resource: $status_code
  operator: equal
  target: 500
  action:
    name: alert_by_email
    parameters:
      - tech_lead@gmail.com
      - critical
```

- Oluwoye will test for each log these rules, because we can have more than one rule for one log entry.
- An example is if someone tries to use a tool and that person is blocklisted by IP addres in your first rule list, and you create a second on for check if the status code of this log will come status `500`. That log will be alert twice, one for the blocklisted person and the second one for status code that said it is internal server error.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
