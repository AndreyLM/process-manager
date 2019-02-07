# PROCESS MANAGER API
===============================================
#### Get process's list
*Returns json data with information about processes*
###### URL
*/v1/process/list*
###### Method
*GET*
###### Success Response:
```javascript
    {
        Code: 200,
        Message: "Success",
        Response: [
            { Name: "Process name", MaxCount: 4, Count: 3},
            { Name: "Process name2", MaxCount: 3, Count: 3},
            ....
        ] 
    }
```
###### Error Response:
```javascript
    {
        Code: 400,
        Message: "Detailed error description",
        Response: "" 
    }
```
###### Sample Call:
```javascript
    $.ajax({
        url: "/v1/process/list",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```
-----------------------------------------------
#### Check if process exist
*Returns json data with information about process existence*
###### URL
*/v1/process/exist/{process_name}*
###### Method
*GET*
###### Success Response:
```javascript
    {
        Code: 200,
        Message: "Success",
        Response: true || false 
    }
```
###### Error Response:
```javascript
    {
        Code: 400,
        Message: "Detailed error description",
        Response: "" 
    }
```
###### Sample Call:
```javascript
    $.ajax({
        url: "/v1/process/exist/processName",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```


    