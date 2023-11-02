pipeline{

    agent any

    stages {

       stage ("build") {
            steps {
                echo 'building the application...'
               
                 sh "docker build -t golangapp:1.3 ."
             }
        }

       stage ("test") {
            steps {
                     echo 'testing the application...'
             }
        }

       stage ("publish") {
            steps {
                       echo ' publish the application...'

                       sh "docker tag golangapp:1.3 eli7890/golangapp:1.3"
                       sh "docker login"
                       sh "docker push eli7890/myapp:1.3"
             }
        }

    }
}
