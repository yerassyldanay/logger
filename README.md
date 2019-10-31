### Logger

This is how I log a message (e.g. an error) into a file. This is not how you should do, but hope it will be helpful for some folks

After running the main.go, you will get logger.log file created (if it does not exists) and following messages in it:
```
DB_ACCESS: 2019/10/31 11:36:40 [TEST_ERROR] Error occured while calling a function. Message: This is an error message. Checking
DB_ACCESS: 2019/10/31 11:36:40 [TEST_ERROR] Operation has been done successfully. Message: This is a success message. Checking
```
