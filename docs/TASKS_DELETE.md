*[content](../README.md)*
### Delete all task 
*Delete all task*
#### URL
*/v1/tasks*
#### Method
*DELETE*
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
        url: "/v1/tasks",
        dataType: "json",
        type: "DELETE",
        success : function(r) {
            console.log(r);
        }
    });
```