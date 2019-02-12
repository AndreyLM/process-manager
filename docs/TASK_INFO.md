*[content](../README.md)*
### Get task info
*Returns json data with information about task*
#### URL
*/v1/task/{taskName}*
#### Method
*GET*
#### URL Params
Required:  
* *taskName=[string]*
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Task info",
        "Response": { 
            "name": "Task name", 
            "maxCount": 4, 
            "data": {
                "param1": value1,
                "Param2": value2,
                ...
            },
            "count": 3
        } 
    }
```
#### Error Response:
```javascript
    {
        "Code": 404,
        "Message": "Detailed error description",
        "Response": null 
    }
```
#### Sample Call:
```javascript
    $.ajax({
        url: "/v1/task/taskName",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```