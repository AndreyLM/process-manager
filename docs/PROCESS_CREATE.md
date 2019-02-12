*[content](../README.md)*
### Create task 
*Create new task*
#### URL
*/v1/task/{taskName}*
#### Method
*POST*
#### URL Params
Required:  
* *taskName=[string]*
#### Data Params
```javascript
    {
        "duration": [integer:required],
        "data": [object:optional]
    }
```
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Process successfully added task",
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