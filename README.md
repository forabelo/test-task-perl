# Test task perl

This README file provides instructions for setting up and running the application. Setup comes in the latest part

## Task related

1. Check the sql directory for the second task in `resources/sql/*`
2. Check the resources/ directory for the task description, the task is in the file `resources/task.pdf`
3. Check the screenshots in `resources/screenshots/*`
4. Check the others documents (like a Postman collection) in `resources/other/*`
5. In the codebase of the project, there are comments with the decisions 

## Setup

1. Copy the `.env.dist` file to a new file named `.env`:

    ```bash
    cp .env.dist .env
    ```

2. Edit the `.env` file and set your environment variables:

    ```bash
    v .env
    ```

3. Save the changes

## Running the Application

1. Build the Docker image. Replace `test-task-perl` with a name for your Docker image:

    ```bash
    docker build -t test-task-perl .
    ```

2. Run the Docker container:

    ```bash
    docker run -p 8441:8441 --env-file .env test-task-perl
    ```
