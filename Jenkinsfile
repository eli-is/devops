pipeline{

    agent any

    stages {

       stage ("build") {
            steps {
                echo 'building the application...'
                sh "ls- la"
                 sh "docker build -t myapp:1.0 ."
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
