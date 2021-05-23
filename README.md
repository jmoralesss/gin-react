# morales-backend-1

### Run the following commands to setup and run
1. `docker pull jmoralesss/client`
2. `docker pull jmoralesss/server`
3. `docker-compose up -d`
4. Visit `https://localhost:3000/` in your browser

## Design Decisions
* Front-End
    * I created a react app using the `npx create-react-app my-app --template typescript` template
* Back-End
    * Gin for the api framework
    * Gorm for the ORM
    * Sqlite for the database
        * When a user creates an account the provided password is hashed using the `golang.org/x/crypto/bcrypt` package and then stored in db

## Dev time was about half a day