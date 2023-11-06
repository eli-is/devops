pipeline {
environment {
registry = "eli7890/myapp"
registryCredential = 'dockerhub_id'
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
