# file-renamer

## Description
In this repository contains scripts for renaming all files within a directory and every nested directory that contain a coma `","` with an underscore `"_"`. Only one script needs to be run depending on which source code is installed on your machine.

## How to Use
If no source code is installed on your machine, install either [Python](https://www.python.org/downloads/) or [Go](https://golang.org/dl/). One installed, complete the following steps:
1. Clone this repository or download the file according the source code installed (Python or Go).
2. Copy either `comaReplace.py` or `comaReplace.go` and paste it into the directory where the files need to be renamed. Note: this only needs to be done once at the parent directory.
3. Using `CMD`, `Terminal`, or `Bash`, navigate to the directory where `comaReplace.*` was copied into. **If you do not know what this is** and you are on a **Windows** machine:
   1. Open the folder you copied the script to.
   2. Locate the address bar of the folder. It should located next to the search bar or back button.
   3. In the address bar, replace all the text with `cmd` and press enter.
   4. A new window should pop. This is Window's command line, `CMD`.
4. Run the script in command line. 
   
   If you have **Python** installed, run: 
   ```
   python comaReplace.py
   ```
   If you have **Go** installed, run:
   ```
   go run comaReplace.go
   ```
5. Every file with a coma within it's file name should now have the coma replaced with an underscore. Verify the renamed files.

## To-do
This script would probably be more accessible if an executable was created. It would also be more powerful (or potentially dangerous) if users could pick which character or characters to be replaced.