docker build . -t barfuss-website  
docker run -d -p 8080:80 barfuss-website
