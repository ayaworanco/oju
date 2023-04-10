use std::io::Result;

use oluwoye::server;

#[tokio::main]
async fn main() -> Result<()> {
    server::start().await.unwrap();
    Ok(())
}
