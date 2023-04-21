# Writing rules

First you need to create a file where you want to run `olu`, and in this folder you need to have a `rules.yaml` file or in a ENV varible called `RULES_YAML_PATH telling the absolute path of your rules.yaml

## Rules example

```yaml
- resource: $ipv4
  operator: equal
  target: 54.36.149.41
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
      - product_team@gmail.com
      - critical
```


- Oluwoye will test for each log these rules, because we can have more than one rule for one log entry.
- An example is if someone tries to use a tool and that person is blocklisted by IP addres in your first rule list, and you create a second on for check if the status code of this log will come status `500`. That log will be alert twice, one for the blocklisted person and the second one for status code that said it is internal server error.

