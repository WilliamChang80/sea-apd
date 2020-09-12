# sea-apd
Project for compfest 12 Software Engineering Academy by APD team

`sea-apd` is an E-Commerce that help customers to buy what they want from top merchants

it's a RESTFUL API that can be used through this documentation [swagger](https://github.com/WilliamChang80/sea-apd)

## Project structure

```sh
sea-apd/
├── common              # Used for utilities and reusable function
      ├── auth          # related to authentication proccess i.e. JWT Token, hash password 
      ├── constant      # const value that will be used (usecase, dto, controller and route) to avoid hardcode value
      ├── mailer        # to handle mail notification
      ├── observer      # listen to notification
├── controller          # for handling outside requestt (like delivery layer)
      ├── http          # Entry point based on HTTP method
├── domain              # Contains declaration of structs and contract for repository, usecase and controller
├── dto                 # Contains decalration of structs
      ├── domain        
      ├── request        
      ├── response
├── infrastucture       # config database
      ├── db
├── mocks               # contains all mocks from any source
      ├── file
      ├── postgres
      ├── repository
      ├── usecase
├── nginx
├── repository          # to handle operation to database
      ├── postgres
├── routes            
├── usecase             # to handle business logic
```

## Features
  - Register Customer, Merchant and Admin
  - Reject Merchant proposal
  - Login
  - Update profile Merchant and Customer
  - CRUD Products
  - Get Merchant products
  - Update Saldo Mechant
  - CRUD transaction (cart)
  - Submit payment
  - Get list requested item
  - Reject transaction
  - View transaction histories
