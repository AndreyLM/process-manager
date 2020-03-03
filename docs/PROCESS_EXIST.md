*[content](../README.md)*
### Check if process exist 
*Returns json data with information about process existence*
#### URL
*/v1/task/{taskName}/{processUUID}/exist*
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
        "Message": "Process exists" || "Process does not exist",
        "Response": true || false
    }
```
#### Error Response:
```javascript
    {
        "Code": 400,
        "Message": "Detailed error description",
        "Response": "" 
    }
```
#### Sample Call:
```javascript
    $.ajax({
        url: "/v1/task/taskName/processUUID/exist",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```