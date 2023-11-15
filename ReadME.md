# Golang Standard Rest-API

    Lorem Ipsum 

# Project Component

    - Modular loosely coupled design
    - Rest API back end service (GIN Framerwork)
    - Event Driven Architecture (Kafka)
    - SQL Database as main datastore (Postgresql)
    - NoSQL Datase as cache and search (Redis,Elasticsearch)
    - CI (Github action)
    - CD (AWS EKS)
    - Project management (JIRA)
    - Logging (Logrus & EFK stack & Docker driver)
    - Database Migration (golang-migrate with rollback mechanism)
    - Application Configuration (JSON file + Env Variable)
    - Validator (Go Playground)
    - Project Repository (Github)
    - Docker Image Registry (Docker Hub)
    - High Availability (Database replication + EKS)
    - Scheduler (gocron)
    - External Integration
    - API documentation (Postman)
    - Automated Testing
      - Jmeter (Load + stress test)
      - Functional Testing (Postman)
      - Regression Testing
      - Testing management tools
      - Unit Testing (Testify)

## Project Structure
- handler >> dilarang ada bisnis logic here , response building, request mapping, 1st layer validator
- db >> all related retrieving data from database handled here. 
- service >> put your bussiness logic here , 2nd layer validator
- Model >> define DTO, DTIN, database model. Create custom method from each model to do specific job such as validation


## Tutorial to create new module/package