# Mini Aspire

Mini Aspire is a simple web application that allows users to create an account, login, and apply for a loan.

### Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes

### Prerequisites

1. Install [Go](https://go.dev/doc/install)
2. Set up Go Environment Variables

Go requires $GOPATH to be set and $GOPATH/bin added to $PATH. This isn't done by default always, so it's best to do it yourself.

In your `~/.bash_profile`, add the following lines at the end of the file :

 ```sh
 export GOPATH=$HOME/go
 export PATH=$PATH:$GOPATH/bin
 ```

If you use zsh as your default shell, add the following line at the end of `~/.zshrc` :

 ```sh
 export PATH=$PATH:/usr/local/go/bin
 ```

3. Check your installation by running `go version` in a new terminal.

4. Set up directories as per convention

   Now we've set up the environment variables, let's actually make those directories.

    ```sh
    mkdir -p $HOME/go/src
    mkdir -p $HOME/go/bin
    ```

   All the Go packages downloaded are kept in either of these 2 folders, depending on the type of package.

5. Install [Docker](https://docs.docker.com/install/)

We are only running mysql in docker to support our application. After installing, you can simply run the following command to start the mysql container:

```sh
docker-compose -f docker-compose.yml up -d --build
```

Required tables will be created automatically when the application is started.

6. Once docker is up and running, we can execute the command ```go run main.go``` to start the application. The application
is automatically configured to connect to the database.