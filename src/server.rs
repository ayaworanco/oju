use std::env;

use tokio::io::AsyncReadExt;
use tokio::net::{TcpListener, TcpStream};

use crate::parser::parse;

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
    // Ok(())
}

async fn handle_stream(stream: &mut TcpStream) {
    loop {
        // FIXME: change this to a dynamic buffer
        let mut buffer = vec![0; 1024];
        let bytes_readed = stream.read(&mut buffer).await.unwrap();
        if bytes_readed == 0 {
            break;
        }

        // parse packet
        let parsed = parse(&mut buffer);
        println!("{:?}", parsed);
    }
}
