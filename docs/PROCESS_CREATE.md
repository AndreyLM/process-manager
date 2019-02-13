*[content](../README.md)*
### Create process 
*Create new process*
#### URL
*/v1/task/{taskName}*
#### Method
*POST*
#### URL Params
Required:  
* *taskName=[string]*
#### Data Params
```javascript
    {
        "duration": [integer:required],
        "data": [object:optional]
    }
```
#### Success Response:
```javascript
    {
        "Code": 200,
        "Message": "Process successfully added to task",
        "Response": {
            "uuid" : "7b33b415-ccb6-46cb-9734-58e918a590de"
        } 
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
        url: "/v1/task/taskName",
        dataType: "json",
        type: "POST",
        data: { 
            "duration": 300, // process life-time in seconds 
            "description" : "description",
            "data" : {
                "offset" : 13,
                "limit": 5
            }
        },
        success : function(r) {
            console.log(r);
        }
    });
```