*[content](../README.md)*
### Delete task 
*Delete task*
#### URL
*/v1/task/{taskName}*
#### Method
*DELETE*
#### URL Params
Required:  
* *taskName=[string]*
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
        url: "/v1/task/taskName",
        dataType: "json",
        type: "DELETE",
        success : function(r) {
            console.log(r);
        }
    });
```