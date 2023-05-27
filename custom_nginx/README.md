# open this directory on your terminal, then run the following command:
  'docker build -t my-nginx .'
# then run this command:
  'docker run -d -p 80:80 -p 443:443 --name my-nginx-container my-nginx'