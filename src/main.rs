use std::io::Read;
use std::net::{TcpListener, TcpStream};
use std::str;

fn main() -> std::io::Result<()> {
    // Creates a listener and binds to 127.0.0.1:7070
    // if something goes wrong will panic, because we
    // does not need an server within an error running
    let listener = TcpListener::bind("127.0.0.1:7070").unwrap();

    // Iterates over all stream incoming to this server
    // stream variable is a TcpStream mutable (because )
    for stream in listener.incoming() {
        // Like an unwrap this will be converted to an Err if there is some problem
        let mut stream = stream?;

        // spawn a new thread and move the stream value to the closure
        std::thread::spawn(move || {
            // if handle client function returns an Err (as if is converted to an Err)
            // this will print an error mesage
            if let Err(e) = handle_client(&mut stream) {
                eprintln!("Error on: {}", e);
            }
        });
    }

    Ok(())
}

// this function need to pass a reference of a mutable TcpStream
// and will return a default result with an empty tuple
fn handle_client(stream: &mut TcpStream) -> std::io::Result<()> {
    let mut buffer = [0; 1024];
    stream.read(&mut buffer)?;
    let readable = str::from_utf8(&mut buffer).unwrap();

    println!("{}", readable);

    Ok(())
}
