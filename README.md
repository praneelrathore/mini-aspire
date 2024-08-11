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


### Notes about the project
- The project is structured in a way that it can be easily extended to include more features. I have exposed interfaces
for the database and service layers, so that we can easily swap out the database or service layer without affecting the
rest of the application.

- Tests have been written for complicated feature like calculating loan amount for each installment appropriately for the number
of terms. I have left the scope of adding tests for trivial cases as well with the interface layer, but due to the interest of time,
I have not added them.

- Some extra features that I have included from my end
  - Added a simple docker-compose file to run mysql in a container
  - User and admins can register themselves. Currently, they are being required to pass through the request payloads, but in real world scenario, they can be identified from cookies.
  - Users can apply for loans. The loan amount is calculated based on the term and the amount requested. Only users who are authenticated can apply for loans.
  - Admins can approve or reject loan requests. The status of the loan is updated in the database. There are validation checks that admin is authorized and only approve loans which need approval.
  - Users can repay their loans in installments. The amount is calculated based on the term and the amount requested. Only users who are authenticated can repay loans.
  - Validations checks are present for user's repayment where there are checks that user is only repaying their own loan, that the loan amount is valid and that the loan is not already repaid.
  - I have also added a check where users will be asked for extra 10% amount in case their due date is passed. This is to ensure that users are incentivized to repay their loans on time.