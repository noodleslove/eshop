### Application Architecture — The Layers

In traditional enterprise applications, it is a common practice to break the application logic into three manageable layers, namely:

- Controller — This is the layer that receives the incoming requests(primarily HTTP), deconstructs the request body, request parameters, headers, path variables etc. to create a model object that the service methods act on. This also is responsible to render the response back to the response stream.
- Service — This is where the business logic resides and interacts with the repository layer.
- Repository — This handles the lower level details of persisting a domain object or structs.

### Defining The Entity Layer

- Customer Struct Type
- Address Struct Type

### Defining The Repository Layer

- CRUD Repository
- Customer Repository
- Address Repository

### Defining The Service Layer

- Customer Service
- Address Service

### Defining The Controller Layer

- Customer Controller
- Address Controller
