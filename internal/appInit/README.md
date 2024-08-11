### AppInit

AppInit is a short form for app initialization. This package is responsible for initializing the application. 
It is the first package to be executed when the application starts. It initializes the application by setting up the configuration, database, and other required services.

Config file helps us read the data from config.yaml. This way, configuration for our application can be simply kept
in a yaml file and can be utilized anywhere in our application.

Database file initializes the connection with the database. It uses the config file to get the database configuration.
In a real world scenario, our application would be running in a containerized environment. In that case, we only
need to initialize the database once and the connection can be kept open until the container is running. For simplicity
purpose in this application, I have initialized the database during api calls.

Env is short form for environment. It is a common practice in golang projects to keep the environment variables in a
variable initialized during the start and then use it throughout the application. This package initializes the environment.