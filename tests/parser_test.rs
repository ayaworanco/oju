#[cfg(test)]
mod parser_test {
    use oluwoye::logger::parser::parse;

    #[test]
    fn should_parse_string_to_log() {
        let test_packet =
            "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO 1.1\n02:49:12\n127.0.0.1 GET / 200";
        let parsed = parse(test_packet).unwrap();
        assert_eq!(parsed.header.verb, "LOG".to_string());
    }
}
