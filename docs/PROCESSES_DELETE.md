*[content](../README.md)*
### Delete all task 
*Delete all processes in task*
#### URL
*/task/{taskName}/clear*
#### Method
*DELETE*
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "All processes were deleted",
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
        url: "/task/taskName/clear",
        dataType: "json",
        type: "DELETE",
        success : function(r) {
            console.log(r);
        }
    });
```