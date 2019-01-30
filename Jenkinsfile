pipeline {
    agent {label 'docker'}

    stages {
        stage('Build') {
            steps {
                sh "docker build -t documentor:latest ."
            }
        }
        stage('Deploy') {
            steps {
                sh "printenv"
                sh "deploying somewhere ...."
            }
        }
    }

    post {
        unsuccessful{
            sh "docker rmi documentor:latest"
        }
    }
}