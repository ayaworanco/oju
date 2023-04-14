#[cfg(test)]
mod rule_test {
    use oluwoye::{logger::parser::parse, ruler::rule::Rule};

    const TEST_PACKET: &str = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO 1.1\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\"";

    const TEST_ERROR_PACKET: &str = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO 1.1\n02:49:12\n192.168.11.11 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\"";
    const RULES_YAML: &str = r#"
- resource: $ipv4
  operator: equal
  target: 54.36.149.41
  action:
    name: alert_by_email
    parameters:
      - tech_lead@gmail.com
      - warning

    "#;

    const RULES_UNKOWN_TARGET_YAML: &str = r#"
- resource: $ipv4
  operator: equal
  target: aaaaaaaaaaaaaaaaaa
  action:
    name: alert_by_email
    parameters:
      - tech_lead@gmail.com
      - warning

    "#;

    #[test]
    fn should_return_true_on_rule() {
        let log = parse(TEST_PACKET).unwrap();
        let rules: Vec<Rule> = serde_yaml::from_str(RULES_YAML).unwrap();
        let rule = &rules[0];
        let result = rule.run(log.message).unwrap();
        assert_eq!(log.header.verb, "LOG");
        assert_eq!(result, true);
    }

    #[test]
    fn should_return_false_on_rule() {
        let log = parse(TEST_ERROR_PACKET).unwrap();
        let rules: Vec<Rule> = serde_yaml::from_str(RULES_YAML).unwrap();
        let rule = &rules[0];

        let result = rule.run(log.message).unwrap();
        assert_eq!(log.header.verb, "LOG");
        assert_eq!(result, false);
    }

    #[test]
    fn should_return_false_on_when_unkown_variable_rule() {
        let log = parse(TEST_PACKET).unwrap();
        let rules: Vec<Rule> = serde_yaml::from_str(RULES_UNKOWN_TARGET_YAML).unwrap();
        let rule = &rules[0];

        let result = rule.run(log.message).unwrap();
        assert_eq!(log.header.verb, "LOG");
        assert_eq!(result, false);
    }
}
