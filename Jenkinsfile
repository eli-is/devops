pipeline{

    agent any

    stages {

       stage ("build") {
            steps {
                echo 'building the application...'
                 sh "docker build -t myapp:1.0 https://github.com/eli-is/devops.git"
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
