*[content](../README.md)*
### Check if task exist 
*Returns json data with information about task existence*
#### URL
*/v1/task/exist/{taskName}*
#### Method
*GET*
#### URL Params
Required:  
* *taskName=[string]*
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Success",
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
        url: "/v1/task/exist/taskName",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```