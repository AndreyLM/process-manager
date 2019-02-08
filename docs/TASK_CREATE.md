*[content](../README.md)*
### Create task 
*Create new task*
#### URL
*/v1/task*
#### Method
*POST*
#### Data Params
```javascript
    {
        Task: {
            Name: [string:required],
            MaxCount: [integer:required],
        }
    }
```
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
        url: "/v1/task",
        dataType: "json",
        type: "POST",
        data: {
            Task: {
                Name: "taskName",
                MaxCount: 2
            }
        },
        success : function(r) {
            console.log(r);
        }
    });
```