use log::{info, warn};
use std::env;
use std::fs::File;
use std::io::Read;

use lazy_static::lazy_static;
use tokio::io::AsyncReadExt;
use tokio::net::{TcpListener, TcpStream};

use crate::logger::parser::parse;
use crate::ruler::rule::Rule;

lazy_static! {
    static ref RULES: Vec<Rule> = {
        let rule_file = env::var("RULES_YAML_PATH").unwrap_or("./rules.yaml".to_owned());
        let mut rules_string = String::new();
        let mut rules_buff = File::open(rule_file).unwrap();
        rules_buff.read_to_string(&mut rules_string).unwrap();

        let rules: Vec<Rule> = serde_yaml::from_str(&rules_string.as_str()).unwrap();
        return rules;
    };
}

pub async fn start() -> std::io::Result<()> {
    let port = env::var("PORT").unwrap_or("8080".to_owned());
    let host = env::var("HOST").unwrap_or("localhost".to_owned());
    println!("[OLUWOYE] started on port {}", port);
    let listener = TcpListener::bind(format!("{}:{}", host, port))
        .await
        .unwrap();

    loop {
        let (mut stream, addr) = listener.accept().await.unwrap();
        println!(
            "[OLUWOYE] accepting new connection: {}",
            addr.ip().to_string(),
        );
        tokio::spawn(async move {
            handle_stream(&mut stream).await;
        });
    }
}

async fn handle_stream(stream: &mut TcpStream) {
    loop {
        // FIXME: change this to a dynamic buffer
        let mut buffer = vec![0; 1024];
        let bytes_readed = stream.read(&mut buffer).await.unwrap();
        if bytes_readed == 0 {
            break;
        }

        let message = std::str::from_utf8(&mut buffer).unwrap();
        match parse(message) {
            Ok(log) => {
                for rule in RULES.iter() {
                    // First here we will log the output
                    println!("something here");
                    info!(
                        "app[{}] time[{}]\n{}",
                        log.header.app_key,
                        log.timer,
                        log.message.clone()
                    );
                    if rule.run(log.message.clone()).unwrap() {
                        warn!("Rule alerted!");
                    }
                }
            }
            Err(error) => println!("{:?}", error),
        }
    }
}
