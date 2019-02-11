*[content](../README.md)*
### Create task 
*Create new task*
#### URL
*/v1/task*
#### Method
*POST*
#### Data Params
```javascript
    {
        name: [string:required],
        maxCount: [integer:optional],
        description: [string:optional],
        data: [object:optional]
    }
```
#### Success Response:
```javascript
    {
        Code: 200,
        Message: "Success",
        Response: "" 
    }
```
#### Error Response:
```javascript
    {
        Code: 400,
        Message: "Detailed error description",
        Response: "" 
    }
```
#### Sample Call:
```javascript
    $.ajax({
        url: "/v1/task",
        dataType: "json",
        type: "POST",
        data: {
            "name": "taskName",
            "description": "some description",
            "maxCount": 2,
            "data" : {
                "param1" : "someValue",
                "param2" : 3
            }
        },
        success : function(r) {
            console.log(r);
        }
    });
```