# Day 3: Connect your backend to a database

Now that you have some routes defined, it is time to connect to a database. You could use a SQL or NoSQL database.

## Setup
NOTE: We'll be using "MySQL" for this setup
- Navigate to the root of this repo.
- Open the 'main.go' file
- Update the following variables with your preferred database details
    - DB_USER
	- DB_PASS
	- DB_DATABASE
- Run the command ```go run ./main.go``` to start the server.
- Visit the url http://127.0.0.1:3000 to check if the connection was successful