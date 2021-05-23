# morales-backend-1

### Run the following commands to setup and run
1. `docker pull jmoralesss/client`
2. `docker pull jmoralesss/server`
3. `docker-compose up -d`
4. Visit `https://localhost:3000/` in your browser

### Design Decisions
* Front-End
    * I created a react app using the `npx create-react-app my-app --template typescript` template
    * The client is signed out once the bearer token is expired
* Back-End
    * Gin for the API framework
        * A bearer token is passed back to the client once the user is authenticated
    * Gorm for the ORM
    * Sqlite for the database
        * When a user creates an account the provided password is hashed using the `golang.org/x/crypto/bcrypt` package and then stored in the db

### Dev time was about half a day