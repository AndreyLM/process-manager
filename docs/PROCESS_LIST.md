*[content](../README.md)*
### Get all processes realted to task 
*Returns json data with information about task's processes*
#### URL
*/v1/task/{taskName}/processes*
#### Method
*GET*
#### Success Response:
```javascript
    {
        Code: 200,
        Message: "Success",
        Response: { 
                Name: "Task name", 
                Processes : [ 
                    { Process: {
                        
                    } } 
                ],
            },
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
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```