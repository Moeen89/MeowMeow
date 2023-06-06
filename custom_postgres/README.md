# open this directory on your terminal, then run the following command:
  ```docker build -t my-postgres-image .```
# then run this command:
  ```docker run -d --name my-postgres-container -p 5432:5432 my-postgres-image```
