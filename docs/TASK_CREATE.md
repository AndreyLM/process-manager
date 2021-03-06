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
        "name": [string:required],
        "maxCount": [integer:optional],
        "description": [string:optional],
        "data": [object:optional]
    }
```
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Task successfully added to collection",
        "Response": null 
    }
```
#### Error Response:
```javascript
    {
        "Code": 400,
        "Message": "Detailed error description",
        "Response": null 
    }
```
#### Sample Call:
```javascript
    $.ajax({
        url: "/v1/task",
        dataType: "json",
        type: "POST",
        data: { 
            "name": "task",
            "maxCount" : 3,
            "description" : "description",
            "data" : {
                "offset" : 13,
                "limit": 5
            }
        },
        success : function(r) {
            console.log(r);
        }
    });
```