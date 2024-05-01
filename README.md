Your task is to create an API that will allow users to rent cars.

Requirements:

- user can rent a car, update the rental date, fetch(get) rental details, and cancel a rental
- user can create an account in the service, see his account details, modify them, and delete his account
- cars can be only rented for a user that has an account in the system

This project must work seamlessly, meaning there will be no runtime errors.

Additionally, I want to be able to run the entire service by typing in my console make server.

Technical requirement:

To create the HTTP server and route paths, projects MUST use this library:
https://github.com/gorilla/mux

 c r u d 
 1. model user 
 2. model rental
 3. store rental,user
 4. rental must have id of user to verify access
 5. handler rental,user
 6. router(mux gorila)
 7. start server 
 8. read documentation again to find what is missing 