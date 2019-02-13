*[content](../README.md)*
### Get process info
*Returns json data with information about process*
#### URL
*/v1/task/{taskName}/{processUUID}*
#### Method
*GET*
#### URL Params
Required:  
* *taskName=[string]*
* *processUUID=[string]*
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Process info",
        "Response": {
            "uuid": "0a9566d0-bd7d-44fa-8175-cd4696ddc617",
            "data": {
                "param1": value1,
                ...
            },
            "status": "",
            "expire": "13/02/2019 09:55:53.4329"
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
        url: "/v1/task/taskName/processUUID",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```