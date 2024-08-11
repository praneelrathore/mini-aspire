### Configs directory

This directory contains the configurations for our application. For example, configuration related to mysql is kept here.
We have also kept other configurations which can be tuned as per the requirement. For example, currently we only have sql related
configurations in that file. But in the future, we might want to add configurations for other services like redis, kafka, logging etc.

Advantages of keeping it this way is that they don't pollute our code logic, can be reused across the application and can be easily
changed without changing the code.