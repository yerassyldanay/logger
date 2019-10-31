### Logger

This is how I log a message (e.g. an error) into a file. This is not how you should do, but hope it will be helpful for some folks

##### Output (run main.go):

```
DB_ACCESS: 2019/10/31 11:36:40 [TEST_ERROR] Error occured while calling a function. Message: This is an error message. Checking
DB_ACCESS: 2019/10/31 11:36:40 [TEST_ERROR] Operation has been done successfully. Message: This is a success message. Checking
```

##### Example #1

```
var error_message = utils.FunctionCallError{utils.FunctionCallSuccess{
    FunctionName: "TEST_ERROR",
    Message:      "This is an error message. Checking",
}}

utils.LogMini(&error_message)
```

##### Example #2

``` 
var success_message = utils.FunctionCallSuccess{
    FunctionName: "TEST_ERROR",
    Message:      "This is a success message. Checking",
}

utils.LogMini(&success_message)
```

##### Note
```
/*
    Close the file at the end of the function
    to avoid a leakage of resources
*/
defer utils.GetLogFile().Close()
```