pipeline {
environment {
registry = "696353136210.dkr.ecr.eu-north-1.amazonaws.com/my-app"
registryCredential = 'dockerhub_idaws'
}
agent any
stages {
    stage('Building our image') {
steps{
script {
dockerImage = docker.build registry + ":2.0"
}
}
}
    stage('Deploy our image') {
steps{
script {
docker.withRegistry( '', registryCredential ) {
dockerImage.push()
}
}
}
}

}
}
