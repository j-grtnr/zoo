# basicapps

This command line tool reads a text file from /basicapps folder and extracts all lines containing a specific keyword. The keyword is returned in red color.  
Command line arguments specify the file, the keyword to search for and whether to ignore case or not.

## Usage

Navigate into /basicapps.
```
go build .
FILE_PATH="hamlet.txt" KEY_STRING="Fortinbras" IGNORE_CASE="true" ./basicapps
```