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
        "message" :  "Email already exists"
    }

```

### Login User - ``POST : /users/login``
- Success
```json

    {
        "status" : "success",
        "message" : "User has been loged in"
    }

```
- Fail
```json

    {
        "success" : false,
        "message" :  "Wrong password or email!"
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
        "message" :  "User does not exist"
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
        "message" :  "Internal server error"
    }

```
