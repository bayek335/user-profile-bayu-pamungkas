## API - User Endpoint


### Register User - ``POST : /users/register``
- Success
```json

    {
        "success" : true,
        "message" : "User has been created",
        "data" : {
            "user_id" : 1,
            "username": "Bayu Pamungkas",
            "email" : "bayupamungkas@gmail.com"
        }
    }

```
- Fail
```json

    {
        "success" : false,
        "message" :  "Email already exists",
        "error_code" : 1306,
    }

```

### Login User - ``POST : /users/login``
- Success
```json

    {
        "status" : "success",
        "message" : "Login sucessfully",
        "data" : {
            "token" : "@token"
        }
    }

```
- Fail
```json

    {
        "success" : false,
        "message" :  "Email or  password incorrect",
        "error_code" : 1306,
    }

```

### Update User - ``PUT : /users/:id``
- Success
```json

    {
        "success" : true,
        "message" : "User has been updated",
        "data" : {
            "user_id" : 1,
            "username": "Bayu Pamungkas Update",
            "email" : "bayupamungkas@gmail.com"
        }
    }

```
- Fail
```json

    {
        "success" : false,
        "message" :  "User does not exist",
        "error_code" : 1306,
    }

```

### Delete User - ``DELETE : /users/:id``
- Success
```json

    {
        "success" : true,
        "message" : "User has been deleted",
        "data" : {
            "user_id" : 1,
            "username": "Bayu Pamungkas Update",
            "email" : "bayupamungkas@gmail.com"
        }
    }

```
- Fail
```json

    {
        "success" : false,
        "message" :  "Internal server error",
        "error_code" : 1306,
    }

```