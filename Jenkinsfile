pipeline{

    agent any

    stages {

       stage ("build") {
            steps {
                  
                sh docker build -t myapp:1.0 .
                echo 'building the application...'
                 
             }
        }

       stage ("test") {
            steps {
                     echo 'testing the application...'
             }
        }

       stage ("deploy") {
            steps {
                       echo 'deploying the application...'
             }
        }

    }
}
