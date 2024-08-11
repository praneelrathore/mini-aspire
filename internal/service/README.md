### Service Layer

Service layer is responsible for implementing the business logic for our application. It sits between the controller
layer and the model/database layer.

Service layer is broken down in different parts according to the different entities of our application. There are 3
primary entities in our application, namely - 
1. User
2. Admin
3. Loan

Apart from it, there is a domains package which contains the domain objects for our application. These domain objects
can be passed around between different layers of our application. 

Each entity has an interface which is being implemented in the entity folder itself. Reason for creating interfaces
are as follows -
1. It allows us to define a contract for the service layer. In future, we might want to change the implementation of
our service layer. In that case, we can create a new implementation of the interface and use that directly. A big
advantage of this is that the controller layer does not need to be changed at all.
2. It allows us to mock the service layer for testing purposes. This is very useful when we want to test the controller.