## Clone from git repository

1. git clone https://github.com/Tanima-062/event-management-service.git [command]

## Dockerize event-management-service app

## Configure Dockerfile

======================================== To Run Docker =============================================
## To build docker and create image

1. docker build -t eventdocker .

## To see docker image 

2. docker ps 

## To run docker 

3. docker run --network="host" -p 8888:8888 -t eventdocker

========================================== TO RUN Local =============================================

## Clone from git repository

1. git clone https://github.com/Tanima-062/event-management-service.git [command]

## Instruction
2. Install Go in PC [If not exist]
3. Go to project direction

## Run this command to start server

4. go run main.go

## All API ENDPOINTS

## Attacched Postman Collection in project directory [event-management-service.postman_collection.json]

1. Event List => http://127.0.0.1:8888/api/events?page=1&limit=2  ["/events"]
   Response => 
    {
        "events": [
            {
                "id": 1,
                "title": "Demo Event",
                "start_at": "2023-11-24T08:48:07+06:00",
                "end_at": "2023-12-15T08:48:07+06:00"
            },
            {
                "id": 2,
                "title": "Demo Event 1",
                "start_at": "2023-12-28T08:48:07+06:00",
                "end_at": "2023-12-29T08:48:07+06:00"
            }
        ],
        "pagination": {
            "total": 6,
            "per_page": 2,
            "total_pages": 3,
            "current_page": 1
        }
    }

2. Event Details => http://127.0.0.1:8888/api/events/1  ["/api/events/:id"] [GET Request]
   Response => 
    {
        "id": 1,
        "title": "Demo Event",
        "start_at": "2023-11-24T08:48:07+06:00",
        "end_at": "2023-12-15T08:48:07+06:00",
        "total_workshops": 2
    }

3. Workshop List => http://127.0.01:8888/api/workshops/1  ["/api/workshops/:eventID"] [GET Request]
   Response => 
    {
        "id": 1,
        "title": "Demo Event",
        "start_at": "2023-11-24T08:48:07+06:00",
        "end_at": "2023-12-15T08:48:07+06:00",
        "workshops": [
            {
                "id": 1,
                "event_id": 1,
                "title": "Demo Workshop",
                "description": "Demo Workshop Description",
                "start_at": "2023-11-24T09:42:12+06:00",
                "end_at": "2023-12-28T09:42:12+06:00"
            },
            {
                "id": 2,
                "event_id": 1,
                "title": "Demo Workshop 1",
                "description": "Demo Workshop 1 description ",
                "start_at": "2023-12-14T09:42:12+06:00",
                "end_at": "2023-12-29T09:42:12+06:00"
            }
        ]
    }

4. Workshop Details => http://127.0.01:8888/api/workshops/detail/1  ["/api/workshops/detail/:id"] [GET Request]
   Response =>
    {
        "id": 1,
        "title": "Demo Workshop",
        "description": "Demo Workshop Description",
        "start_at": "2023-11-24T09:42:12+06:00",
        "end_at": "2023-12-28T09:42:12+06:00",
        "total_reservations": 3
    }

5. Create Reservation => http://127.0.0.1:8888/api/reservation/create/1 ["/api/reservation/create/:workshopID"] [POST Request]

   =======================Case 1=======================================
   
   Request Body => 
   {
	"name": "Famina Ishana",
	"email": "famina@gmail.com"
   }

   Response Body => 
    {
        "reservation": {
            "id": 17,
            "name": "Famina Ishana",
            "email": "famina@gmail.com"
        },
        "event": {
            "id": 1,
            "title": "Demo Event",
            "start_at": "2023-11-24T08:48:07+06:00",
            "end_at": "2023-12-15T08:48:07+06:00"
        },
        "workshop": {
            "id": 1,
            "title": "Demo Workshop",
            "description": "Demo Workshop Description",
            "start_at": "2023-11-24T09:42:12+06:00",
            "end_at": "2023-12-28T09:42:12+06:00"
        }
    }

    =============================Case 2========================
    
    Request Body =>

    {
        "name": "",
        "email": ""
    }

    Response Body =>

    {
        "error": "Key: 'SaveInput.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'SaveInput.Email' Error:Field validation for 'Email' failed on the 'required' tag"
    }