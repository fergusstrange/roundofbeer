[![CircleCI](https://circleci.com/gh/fergusstrange/roundofbeer.svg?style=shield)](https://circleci.com/gh/fergusstrange/roundofbeer) [![Go Report Card](https://goreportcard.com/badge/github.com/fergusstrange/roundofbeer)](https://goreportcard.com/report/github.com/fergusstrange/roundofbeer) [![Coverage Status](https://coveralls.io/repos/github/fergusstrange/roundofbeer/badge.svg)](https://coveralls.io/github/fergusstrange/roundofbeer)

# Round of Beer

Keep track of who's round it is and share it with your friends...

But really this is an example of end to end code, testing and deployment to cloud infrastructure.

### Code

* React front end using Material-UI framework
* Go backend using Gin http web framework
* DynamoDB persistence in AWS  

### Testing

* Front end behaviour testing using TestCafe
* Front end unit test coverage using jest and snapshot rendering
* Backend unit testing using testify
* Backend behaviour testing using [apitest](https://apitest.dev)
* Consumer Driven Contract tests using Pact Framework in both Javascript and Golang
* CircleCI for continuous integration
* Go Report Card
* Coveralls coverage analysis

### Deployment
* Apex up for serverless deployment to AWS Lambda
