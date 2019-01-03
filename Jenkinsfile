pipeline {
    agent any 

    stages {
        stage('Build') {
            steps {
                docker build -t documentor-image .
            }
        }
    }
}