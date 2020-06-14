# gowebservice-heroku
This git repository is forked from jaya-p/gowebservice. The purpose is for adaptation for deploying to heroku.com platform.  

## Deploy to Heroku (Using Heroku CLI)
heroku login  
heroku container:push web -a myexip  
heroku container:release web -a myexip  
heroku logs --tail -a myexip

## Access deployed application
https://myexip.herokuapp.com/  
https://myexip.herokuapp.com/api/v1/helloworld
