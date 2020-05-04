Mimic Unix grep Utility.  
It searches many lines of text, typically piped in from another command, for a given string.

* make a generic function that takes `target` a strig to search and a **BufRead** that could be a file or stdin
* collect the args and put all files into  `Vec<PathBuf>`
* iterate over the Vec and return lines that contain the `target`