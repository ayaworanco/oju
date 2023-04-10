use super::log::{HeaderBuilder, Log, LogError};

pub fn parse(packet: &str) -> Result<Log, LogError> {
    let clean_string = get_clean_string(packet);
    match parse_to_log(clean_string) {
        Ok(log) => Ok(log),
        Err(error) => Err(error),
    }
}

fn get_clean_string(packet: &str) -> String {
    packet
        .trim_matches(char::from(0))
        //.replace("\n", "")
        .to_string()
}

fn parse_to_log(clean_string: String) -> Result<Log, LogError> {
    let parts = clean_string.lines().collect::<Vec<&str>>();
    let (header, timer, message) = (parts[0], parts[1], parts[2]);

    match HeaderBuilder::new(header.to_owned()).build() {
        Ok(header) => Ok(Log {
            header,
            timer: timer.to_owned(),
            message: message.to_owned(),
        }),
        Err(err) => Err(err),
    }
}
