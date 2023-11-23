## API - User Endpoint


### Getting All Images by User - ``GET : /photos``
- Success
```json

    {
        "success": true,
        "message": "Photo succesfully taken",
        "data": [
            {
                "id": 9,
                "title": "Photo 2 already posted guys",
                "caption": "Photo 2",
                "photo_url": "localhost:8080/public/images/jEYldhaBRccohRe2.jpg",
                "user_id": 2
            },
            {
                "id": 8,
                "title": "The transparent glasses after edited",
                "caption": "It is a barbarian picture guys",
                "photo_url": "localhost:8080/public/images/lnRMSdGkNqmmYCH2.png",
                "user_id": 2
            }
        ]
    }

```

### Upload Image - ``POST : /photos``
- Success
```json
    {
        "success": true,
        "message": "Photo succesfully created",
        "data": 
        {
            "id": 9,
            "title": "Photo 2 already posted guys",
            "caption": "Photo 2",
            "photo_url": "localhost:8080/public/images/jEYldhaBRccohRe2.jpg",
            "user_id": 2
        }
    }

```
- Fail
```json

    {
        "succes": false,
        "message": "File must be type of image '.jpg', '.jpeg', '.png' !"
    }

```

### Update Image - ``PUT : /photos/:id``
- Success
```json

    {
        "success": true,
        "message": "Photo succesfully updated",
        "data": {
            "id": 8,
            "title": "The transparent glasses after after edited",
            "caption": "Sorry guys my barbarian has been tired",
            "photo_url": "localhost:8080/public/images/eeQrCUSnQxhLkWW2.png",
            "user_id": 2
        }
    }

```
- Fail
```json

    {
        "succes": false,
        "message": "Photos does not exist!"
    }

```

### Delete Image - ``DELETE : /photos/:id``
- Success
```json

    {
        "success": true,
        "message": "Photo succesfully deleted",
        "data": {
            "id": 8,
            "title": "The transparent glasses after after edited",
            "caption": "Sorry guys my barbarian has been tired",
            "photo_url": "localhost:8080/public/images/eeQrCUSnQxhLkWW2.png",
            "user_id": 2
        }
    }

```
- Fail
```json

    {
        "succes": false,
        "message": "record not found"
    }

```
