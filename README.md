# status_check 
[![Go Report Card](https://goreportcard.com/badge/github.com/samlotti/status_check)](https://goreportcard.com/report/github.com/samlotti/status_check)
[![Go Reference](https://pkg.go.dev/badge/github.com/samlotti/status_check.svg)](https://pkg.go.dev/github.com/samlotti/status_check)
Easy to deploy and configure server status page.  One executable and a configuration file.

This page can be accessed via the browse or ping monitoring services such as statuscake.com / pingdom.com.  Has helped me add monitoring on some smaller servers that primarily run as backend machines. 
The server will inspect the checks and report on the pass or fail.  If any fail then the response will be http 500 with the list of passed / failed tests.

If all are ok, the response is 200 with a list of all the tests.

The configuration is default named "server_check.conf" or the name passed in on the command line.

The configuration is a simple text file. Ex:

config file is typeOfCheck:nameOfCheck:parameters for check
```
## Listen on this port
port:23323
##
## Make sure I have a process running, mongod in this case
## "process" command
##
process:MyTestProcess:mongod
##
## Check backup files are being generated.
## Make sure these 2 files are available and not older than 1d 6h
## Can also check min file size as well.
## Note the last entry are key=value pairs for the "file" command
##
file:habackup:/home/bk/hagames/hagamesdb.sql.gz:maxAge=P1DT6H,minSize=5000
file:ctbackup:/home/bk/backup/current.sql.gz:maxAge=P1DT6H
##

```

Step to deploy
* ./buildLinux.sh
* scp to the server
* create a config file
* Start it with 
./server_check  
* add a crontab entry is desired to keep it running and to restart if server reboots
```cron
* * * * * /home/bk/server_check/server_check  > /home/bk/server_check/cron.out  2>&1
```

Example result page:

    Server status check, v0.1.2
    
    --------------------------------------------------------------------------------
    (   )  ->  MyTestProcess              note: pid: 5385
    (   )  ->  knbackup                   note: size:10.6 GB, age:3h53m14.068808133s
    (   )  ->  gtbackup                   note: size:3.8 GB, age:7h37m28.509473752s
    --------------------------------------------------------------------------------
    Status: OK

The process command will show the pid and the file commands will show the size and age.
The process name and file names are not shown and instead the test description is shown.


Currently, the tests available are:
* process - check the existence of a process with the exec name
* file - check existence of a file and additional make sure min size and max age are met (optional).

more will be added as needed.

Contributions are welcome for more tests.  Creating a test is fairly straightforward.  
* Implement IRule  (see file_checker and process_checker)
* Change LoadConfig to accept your rule name
* Create a parser for your rule (parameters are seperated by : )

Ex:
file:{testname}:{file parameters}:others:others...



