## mobydev-task
* `mkabyken`


### Description:
Test project with function login,register,logout and edit user.


### Usage
Clone the repository:
```
git clone https://github.com/MirasDragonite/mobydev-task.git
```


####   Run a program:


```
go run ./cmd/main.go
```
Open the link in browser
```
localhost:8000
```
 

## Documentation

```
POST `localhost:8000/auth/sign-up` to create new user.
{
    "email": "example@email.com",
    "password": "12345678"
}

POST `localhost:8000/auth/sign-in` to login into system.
{
    "email": "example@email.com",
    "password": "12345678",
    "repeatPassword": "12345678"
}

POST `localhost:8000/auth/logout` to logout from the system.

GET `localhost:8000/edit?id='YourID'` to get data from user edit page

POST `localhost:8000/edit?id='YourID'` to edit user data
{
    "username": "newUserName",
    "email":"newexample@email.com",
    "phone_number": "+777777777",
    "birth_date": "02 January 2006"
}
```