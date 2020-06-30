use std::fs;
use std::io;
use std::path::Path;
use std::error::Error;

/// Copy the existing directory `src` to the target path `dst`.
fn copy_dir_to(src: &Path, dst: &Path) -> io::Result<()> {
    if !dst.is_dir() {
        fs::create_dir(dst)?;
    }

    for entry_result in src.read_dir()? {
        let entry = entry_result?;
        let file_type = entry.file_type()?;
        copy_to(&entry.path(), &file_type, &dst.join(entry.file_name()))?;
    }

    Ok(())
}

/// Copy whatever is at `src` to the target path `dst`.
fn copy_to(src: &Path, src_type: &fs::FileType, dst: &Path) -> io::Result<()>{
    if src_type.is_file() {
        fs::copy(src, dst)?;
    } else if src_type.is_dir() {
        copy_dir_to(src, dst)?;
    } else {
        return Err(io::Error::new(io::ErrorKind::Other, format!("don't know how to copy: {}", src.display())));
    }
    Ok(())
}


fn main() -> Result<(), Box<dyn Error>> {
    let mut args = std::env::args().skip(1);
    let src_dir = match args.next() {
        Some(s) => s,
        None => Err("usage: grep SRC_DIR DST_DIR")?
    };
    let dst_dir = match args.next() {
        Some(s) => s,
        None => Err("usage: grep SRC_DIR DST_DIR")?
    };
    copy_dir_to(Path::new(&src_dir), Path::new(&dst_dir))?;

    Ok(())
}
