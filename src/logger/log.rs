#[derive(Debug)]
pub struct Log {
    pub header: Header,
    pub timer: String,
    pub message: String,
}

#[derive(Debug)]
pub struct Header {
    pub verb: String,
    pub app_key: String,
    pub version: String,
}

pub struct HeaderBuilder {
    pub verb: String,
    pub app_key: String,
    pub version: String,
}

#[derive(Debug)]
pub enum LogError {
    ParseLogError(String),
}

impl HeaderBuilder {
    pub fn new(header: String) -> Self {
        // split this into correct header
        let parts: Vec<&str> = header.as_str().split(" ").collect();
        let (verb, app_key, version) = (parts[0], parts[1], parts[2]);
        HeaderBuilder {
            verb: verb.to_owned(),
            app_key: app_key.to_owned(),
            version: version.to_owned(),
        }
    }

    pub fn build(&self) -> Result<Header, LogError> {
        if self.verb == "".to_string()
            || self.app_key == "".to_string()
            || self.version == "".to_string()
        {
            Err(LogError::ParseLogError(
                "Header must be fullfilled".to_owned(),
            ))
        } else {
            Ok(Header {
                verb: self.verb.clone(),
                app_key: self.app_key.clone(),
                version: self.version.clone(),
            })
        }
    }
}
