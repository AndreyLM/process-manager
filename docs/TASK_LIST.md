*[content](../README.md)*
### Get task's list 
*Returns json data with information about tasks*
#### URL
*/v1/tasks*
#### Method
*GET*
#### Success Response:
```javascript
    {
        Code: 200,
        Message: "Success",
        Response: [
            { Name: "Task name", MaxCount: 4, Count: 3},
            { Name: "Task name2", MaxCount: 3, Count: 3},
            ....
        ] 
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