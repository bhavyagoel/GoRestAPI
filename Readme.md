
![Logo](https://miro.medium.com/max/920/1*CdjOgfolLt_GNJYBzI-1QQ.jpeg)

# Golang RESTful API for Instagram
A Go based REST API built using native libraries. The API has been thoroughly worked through with Postman.

### Routes include 
- **User Route**: These route includes the functionality of ```fetchUser``` and ```addUser```.
- **Auth Route**: Authentication for ```user``` is done using ```jwt```.
- **FileHandler Route**: Worked with ```GridFS``` to store files as ```fs.chunks``` . Added functionality to ```uploadFile``` and ```downloadFile```. 
- **Post Route**: Functionality to create a post having functionality to ```uploadFile, desctiption, title```.

![MongoDB](https://i0.wp.com/www.differencebetween.com/wp-content/uploads/2018/01/Difference-Between-Firebase-and-MongoDB-fig-1.png?resize=30%2C30&ssl=1) ```MongoDB``` is serving as the required database for all the storage functionality of this api. 
To access the **MongoDB** locally we used a docker container to run an instance of it binding the same to localhost port 27017 using ```-p 27017:27017```.
```bash
docker run --name MongoREST -p 27017:27017 -d mongo:latest
```
### Run Locally
Clone the project and navigate to the cloned repo. 
```bash
  git clone https://github.com/bhavyagoel/GoRestAPI.git
  cd GoRestAPI
```
Install go dependencies by initialising a ```local env```
```bash
go mod init
go mod tidy
```
Run the API Server
```bash
go run main.go
```
### API Reference

#### Get details by ```UserID```
```http
  GET /user/getUserById
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required** UserID to fetch details |

**Dummy Response**
```json
{
    "Email": "jane_doe@testing.com",
    "Password": "$2a$14$vJSwx9sOr2e9KMLiVji9tOfB5AbMjro69R/D1wN5Yqa/IpFMCr2Tq",
    "_id": "61617e08b11882299ca8a3fa",
    "id": "393d0be9-1091-4bd0-822b-915980d94b6d",
    "name": "Jane Dow",
    "post": [
        "56f18233-07bc-4472-9f71-4192f3cc2a3f.jpg",
        "927d4eef-7a77-4865-9cb3-0907fa44a871.jpg"
    ]
}
````

#### Add User To Database
```http
  POST /user/addUser
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Name ` | `string` | Your Name as query Parameter |
| `Email ` | `string` | Your Email as query Parameter |
| `Password ` | `string` | Your Password as query Parameter. *Hashed of the same will be saved in db* |

Note: The string password from the ```POST``` request will be automatically hashed using ```golang.org/x/crypto/bcrypt```.  
**Dummy Response**
```json
{
    "message": "User added successfully",
    "result": {
        "InsertedID": "6161e51740c7852d0e45a540"
    }
}
```

#### Adding a ```POST``` by ```UserId```
To add A post Under a user with a given Id following are the query  parma 

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Id ` | `string` | UUID Generated By ```github.com/google/uuid``` package |
| `Title ` | `string` | Title to Your Post passed as a String |
| `Description` | `string` | A paragraph long Description of the post |
| `urlToImage` | `string` | URI to the local image location |
| `publishedAt` | `string` |TimeStamp Generated using ```time.now()```|
| `fileName` | `string` | Unique fileName Generated ```github.com/google/uuid``` |
| `userId` | `string`  | ```userId``` that links the post and the user|

## Testing

To run tests, run the following command
The unit tests can be found nested inside the test directory

```bash
    cd test/
    go test -v 
```

  
## TODO 
- [ ]  Containerising the API using docker 
- [ ] Deploying the API to Postman 

  
## Badges
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)

  
## Creator
- [@Bhavya Goel](https://www.github.com/bhavyagoel)

  
## Screenshots
Screenshots for the same can be found under the ```img``` sub-folder.

  