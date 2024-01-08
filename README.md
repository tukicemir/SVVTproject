# SVVTproject

This is `klix.ba` project for SVVT course. To use it, you need to download `chromedriver` first and place it 
in this directory.

# Installation and usage

```bash
$ go build
$ ./svvt_project --help
Usage of ./svvt_project:
  -headless
    	start chrome in headless mode
  -verbose
    	use verbose mode (print error messages)
$ ./svvt_project -verbose
#1 Register (Regression) [PASS]
#2 Login (Regression) [PASS]
#3 Search (Regression) [PASS]
#4 Article Navigation (Regression) [PASS]
#5 Comment (Regression) [FAILED] comment limit
#6 Categories (Regression) [PASS]
#7 Tags (Regression) [PASS]
#8 Jobs Listing (Regression) [PASS]
#9 Jobs search (Regression) [PASS]
#10 Jobs Apply (Regression) [PASS]
#11 Load test (Smoke) [PASS]
Total tests: 11
Passed tests: 90.91%
Failed tests: 9.09%
```