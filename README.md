# <p align="center">PROCESS MANAGER API</p>
---------------------------------------------
## <a name="content"><p align="center">Content</p>
* [Get process's list](#processList)
* [Check if process exist](#processExist)
* [Get process info](#processInfo)
* [Create new process](#processCreate)
---------------------------------------------
#### <a name="processList"></a>Get process's list *[content](#content)*
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
#### <a name="processExist">Check if process exist *[content](#content)*
*Returns json data with information about process existence*
###### URL
*/v1/process/exist/{processName}*
###### Method
*GET*
###### URL Params
Required:  
* *processName=[string]*
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
-----------------------------------------------
#### <a name="processInfo">Get process info *[content](#content)*
*Returns json data with information about process*
###### URL
*/v1/process/info/{process_name}*
###### Method
*GET*
###### Success Response:
```javascript
    {
        Code: 200,
        Message: "Success",
        Response: { Name: "Process name", MaxCount: 4, Count: 3} 
    }
```
###### Error Response:
```javascript
    {
        Code: 404,
        Message: "Detailed error description",
        Response: "" 
    }
```
###### Sample Call:
```javascript
    $.ajax({
        url: "/v1/process/info/processName",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```
------------------------------

#### <a name="processCreate">Create process *[content](#content)*
*Create new process*
###### URL
*/v1/process*
###### Method
*POST*
###### Success Response:
```javascript
    {
        Code: 200,
        Message: "Success",
        Response: 
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
        url: "/v1/process/info/processName",
        dataType: "json",
        type : "GET",
        success : function(r) {
            console.log(r);
        }
    });
```


    