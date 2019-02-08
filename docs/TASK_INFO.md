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
        Code: 200,
        Message: "Success",
        Response: { 
            Name: "Task name", 
            MaxCount: 4, 
            Count: 3,
            Data: {
                Param1: value1,
                Param2: value2,
                ...
            }
        } 
    }
```
#### Error Response:
```javascript
    {
        Code: 404,
        Message: "Detailed error description",
        Response: "" 
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