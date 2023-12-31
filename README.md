# Teacher API

### Instructions to run with Docker

1. Clone this repository.
2. With the repository as the working directory, run the following command:
   - `docker compose up`
3. The API will be accessible on `localhost:8080`

### Postman API
The collection of example requests for the various endpoints can be accessed [at this link](https://www.postman.com/jamesyeap/workspace/govtech-onecv-internship-assessment/collection/16444163-4a80a196-d727-40d9-bbe0-8ec5e60216d4?action=share&creator=16444163).

### Notes
- If a request references a non-existent student (i.e. register `studentmary@gmail.com` to some student, which does not yet exist in the database), 
  a new record for the student will be created in the database.
- Likewise, if a request references a non-existent teacher (i.e. register some student(s) to `teacherken@gmail.com`, which does not yet exist in the database),
  a new record for the teacher will be created in the database.
