# Simplebank

This project using Go 1.21 and PostgreSQL 14 to create a simple bank. A simple bank involves creating users, bank accounts with balances and transaction histories associated with users, and lists of historical transactions.
This project also utilizes AWS by containerizing the application into a Docker container and storing this container on AWS EC2. I've setup github workflows to automatically upload changes to the container to the production container once the necessary tests pass.

Setup:
You will need to install Go 1.21, PostgreSQL, Docker Desktop locally, as well as setup necesary AWS features such as RDS, EC2, Secrets Manager to deploy a production instance, but for demonstration, I only recommend local showcasing.

Once the repo is pulled and necessary features are installed, use the routers built into the application with the desktop postman to load data into the database, making sure to account for SQL restraints, then you can use the routers to manipulate this data, ensuring ACID transactions and expected flow.\

Check out the GoMock test cases built, including the percentage of code coverage and edge cases considered.

Please email me with any questions, I can go deeper into detail about setup and execution: jordanweeks123.jw@gmail.com.
