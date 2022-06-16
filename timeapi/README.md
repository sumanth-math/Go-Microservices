**"SomeTime Inc" wants to build an API that returns the current time**

Below are the requirements

    The endpoint should be exposed as /api/time

    The output should be in JSON

    {
        "current_time": "2021-08-09 11:18:06 +0000 UTC"
    }

    A user can also request to return the current time in another timezone. /api/time?tz=America/New_York

  A list of TZ database names are available at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

    If thetzparameter is not provided in the URL, it should return the time in UTC

    In case of an invalid TZ database, the API should return the error message “invalid timezone” along with the status code 404

    Bonus: A user can request time in multiple timezones i.e. /api/time?tz=America/New_York,Asia/Kolkata

API Response:

    {
        "Asia/Kolkata": "2021-08-09 01:23:42 +0530 IST",
        "America/New_York": "2021-08-09 01:23:42 -0400 EDT"
    }


The below program will help you with the timezone conversion

    package main
     
    import (
        "fmt"
        "time"
    )
     
    func main() {
        loc, _ := time.LoadLocation("Asia/Shanghai")
        fmt.Println(time.Now().In(loc))
     
        loc, _ := time.LoadLocation("America/New_York")
        fmt.Println(time.Now().In(loc))
    }


Questions for this assignment

Build the time API based on the instructions and the requirements and handle errors appropriately

**How to Run the Program**
1) Open a terminal and go to the timeapi folder.
2) run the command "go run main.go"
3) Open Postman
4) Create a new request in Postman to connect to 
   a)  http://localhost:8090
    This should return a response shown below
    {
    "current_time": "2022-06-16 11:27:00.8879267 +0000 UTC"
    }
    b) http://localhost:8090/api/time?tz=America/New_York,Asia/Shanghai,Asia/Kolkata
    This should return a response shown below
    {
    "America/New_York": "2022-06-16 08:31:15.1744962 -0400 EDT",
    "Asia/Kolkata": "2022-06-16 18:01:15.1760573 +0530 IST",
    "Asia/Shanghai": "2022-06-16 20:31:15.1755334 +0800 CST"
    }
    c) http://localhost:8090/api/time?tz=America/NY,Asia/Shanghai,Asia/Kolkata
    This should return a response shown below
    "invalid timezone America/NY"
