### Deployment

The deployment directory contains the migrations that are needed to setup the database tables for our application.
These SQL migrations are currently configured to run during MySQL docker container initialization. This is done for 
simplicity as of now. In real world scenarios, we can use something like [flyway](https://github.com/flyway/flyway) to 
manage our migrations.