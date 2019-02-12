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
            "description": "",
            "processes": [
                {
                    "uuid": "2c089277-ff0c-4f9f-9806-b0cc99a5da5d",
                    "data": null,
                    "status": "",
                    "expire": "12/02/2019 11:40:42.9607"
                },
                .....
            ],
            "data": {
                "param1": value1,
                "Param2": value2,
                ...
            }, 
            "maxCount": 4, 
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