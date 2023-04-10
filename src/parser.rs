pub fn parse(buffer: &mut Vec<u8>) -> String {
    get_clean_string(buffer)
}

fn get_clean_string(buffer: &mut Vec<u8>) -> String {
    let message = std::str::from_utf8(buffer).unwrap();
    message
        .trim_matches(char::from(0))
        .replace("\n", "")
        .to_string()
}
