### Model Layer

The model layer is responsible for managing the data of the application. It sits between the service layer (which
holds the business logics) and the database. 

The layer has an interface IDatabase. Reason for using interface here are as follows - 
1. It allows us to define a contract for the database layer. In future, we might not want to use MySQL for our database
but instead some other DB like a NoSQL one. In that case, we can create a new implementation of IDatabase and use that
directly. A big advantage of this is that the service layer does not need to be changed at all.
2. It allows us to mock the database layer for testing purposes. This is very useful when we want to test the service layer

This layer can also have a caching layer if needed. I have not included it in current implementation to keep things
simple.