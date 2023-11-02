pipeline{

    agent any

    stages {

       stage ("build") {
            steps {
                echo 'building the application...'
               
                 sh "docker build -t myapp:1.3 ."
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
                       sh"docker tag myapp:1.3 eli7890/myapp:1.2"
                       sh "docker push eli7890/myapp:1.3"
             }
        }

    }
}
