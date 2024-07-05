# Attendance-Management-System

In this project, I have created an attendance management system using golang and mongo, hosted by Docker. It provides the features creating a new id for a student, getting all student details from the id, enter the attendance of a specific student for a specific date, getting the attendance of a particular subject for all dates, getting the total attendance of a student. 

# Pre-requisites
1. Golang
2. Docker
3. MongoDB

# Project Structure
1. main.go :  Contains the main application code
2. docker-compose.yaml : Configuration file for docker compose for mongoDB
3. infra.go : Infrastructure set up of MongoDB
4. model.go : Data models
5. router.go : Router configuration, contains API paths
6. handler.go : Handles requests from router
7. service.go : Provides Business logic for handler


