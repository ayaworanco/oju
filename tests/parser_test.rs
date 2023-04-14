#[cfg(test)]
mod parser_test {
    use oluwoye::logger::parser::parse;

    const TEST_PACKET: &str = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO 1.1\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\"";

    #[test]
    fn should_parse_string_to_log() {
        let parsed = parse(TEST_PACKET).unwrap();
        assert_eq!(parsed.header.verb, "LOG".to_string());
    }
}