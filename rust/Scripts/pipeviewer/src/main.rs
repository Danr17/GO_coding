pub mod args;
pub mod read;
pub mod stats;
pub mod write;

use crossbeam::channel::{bounded, unbounded};
use std::io::Result;
use std::thread;

use args::Args;
use read::read_loop;
use stats::stats_loop;
use write::write_loop;

fn main() -> Result<()> {
    let args = Args::parse();
    let Args {
        infile,
        outfile,
        silent,
    } = args;

    let (stats_tx, stats_rx) = unbounded();
    let (write_tx, write_rx) = bounded(1024);

    let read_handle = thread::spawn(move || read_loop(&infile, stats_tx, write_tx));
    let stats_handle = thread::spawn(move || stats_loop(silent, stats_rx));
    let write_handle = thread::spawn(move || write_loop(&outfile, write_rx));

    let read_io_result = read_handle.join().unwrap();
    let stats_io_result = stats_handle.join().unwrap();
    let write_io_result = write_handle.join().unwrap();

    read_io_result?;
    stats_io_result?;
    write_io_result?;

    Ok(())
}
