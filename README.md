# shorten-url API in Go with Fiber, GORM and Docker

This repository is a demo implementation of Tiny url  

## Pre-requisite
- Docker Engine
- Go

## Docker
Clone this repository and run under apps folder:
```
docker-compose up
```

You can then hit the following endpoints:

| Method | Route                              | Body                                              |
| ------ | ---------------------------------- | --------------------------------------------------|
| POST   | http://localhost:3000/urls/shorten |`{"url": "https://www.amazon.com/", "email": "bob-anderson@gmail.com"}`|
| GET    | http://localhost:3000/urls/list    |  userid=1  List all the links created by this user|
| DELETE | http://localhost:3000/urls/delete  |  id=1                                             |
| GET    | http://localhost:3000/urls/domain  |  userid=1 domain hit count                        |


After shorten API called a similar link will come in response
If clicked on short url will redirected to original url 

```
http://localhost:3000/184d6db5
```

## More Modification are awaited
