# gin-react

### Run the following commands to setup and run
1. Pull down repo
2. Go to the root of repo
3. Run `docker-compose up -d`
4. Visit `https://localhost:8080/` in your browser

### Design Decisions
* Docker
    * The client is served using nginx
* Front-End
    * I created a react app using the `npx create-react-app my-app --template typescript` template
* Back-End
    * Gin for the API framework
        * A bearer token is passed back to the client once the user is authenticated
    * Gorm for the ORM
    * Sqlite for the database
        * When a user creates an account the provided password is hashed using the `golang.org/x/crypto/bcrypt` package and then stored in the db

### Dev time was about half a day
