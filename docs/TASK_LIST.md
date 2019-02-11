*[content](../README.md)*
### Get task's list 
*Returns json data with information about tasks*
#### URL
*/v1/tasks*
#### Method
*GET*
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Tasks list",
        "Response": [
            { 
                "name": "Task name", 
                "maxCount": 4, 
                "count": 3,
                "data": {
                    "param1": "value1",
                    "param2": "value2",
                    ....
                }
            },
            { 
                "name": "Task name2", 
                "maxCount": 3, 
                "count": 3,
                "data": {
                    "param1": "value1",
                    "param2": "value2",
                    ....
                }
            },
            ....
        ] 
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
        url: "/v1/tasks",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```