# PROCESS MANAGER API
===============================================

#### Get process's list
*Returns json data with information about processes*
###### URL
**/v1/process/list**
###### Method
**GET**
###### Success Response:
* **Code: 200**
* **Message: "Success"**
* **Response: { id : 12, name : "Michael Bloom" }**
###### Error Response:
* **Code: 400**
* **Message: "Detailed error description"**
* **Response: {}**
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


    