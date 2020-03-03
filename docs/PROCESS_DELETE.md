*[content](../README.md)*
### Delete process 
*Delete process*
#### URL
*/v1/task/{taskName}/{processUUID}*
#### Method
*DELETE*
#### URL Params
Required:  
* *taskName=[string]*
* *processUUID=[string]*
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Process was deleted",
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
        url: "/v1/task/taskName/processUUID",
        dataType: "json",
        type: "DELETE",
        success : function(r) {
            console.log(r);
        }
    });
```