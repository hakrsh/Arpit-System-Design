# MySQL + Node.js + PHPMyAdmin Demo

This demo illustrates a setup using Docker Compose to run a MySQL database, a Node.js application, and PHPMyAdmin. The Node.js application inserts dummy users into the MySQL database.

## Getting Started

### Prerequisites

Make sure you have Docker and Docker Compose installed on your machine.

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

### Running the Demo

1. Clone the repository:

    ```bash
    git clone git@github.com:harikrishnanum/Arpit-System-Design.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Arpit-System-Design-for-Beginners/relational_database
    ```

3. Run Docker Compose to start the containers:

    ```bash
    docker-compose up -d
    ```

4. Open PHPMyAdmin in your browser:

    - URL: [http://localhost:8001](http://localhost:8001)
    - Credentials:
      - **Username:** root
      - **Password:** root

5. Access the Node.js container:

    ```bash
    docker exec -it relational_database-node-1 bash
    ```

6. Navigate to the Node.js app directory:

    ```bash
    cd app
    ```

7. Install Node.js dependencies:

    ```bash
    npm install
    ```

8. Run the Node.js script to insert dummy users:

    ```bash
    node insert_dummy_users.js
    ```

## Additional Information

- **MySQL Database:**
  - Host: mysql
  - Username: root
  - Password: root
  - Database: social_network

- **Node.js Application:**
  - The Node.js application is located in the `app` directory.
  - The script `insert_dummy_users.js` inserts dummy users into the MySQL database.

- **PHPMyAdmin:**
  - PHPMyAdmin is accessible at [http://localhost:8001](http://localhost:8001).
  - Use the credentials: Username: root, Password: root.

## Notes

- The demo uses Docker Compose to orchestrate the MySQL, Node.js, and PHPMyAdmin containers.
- PHPMyAdmin is used for easy MySQL database management.
- The Node.js application inserts dummy users into the MySQL database.
