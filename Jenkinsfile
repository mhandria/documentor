pipeline {
    agent {label 'docker'}

    stages {
        stage('Build') {
            steps {
                sh "docker build -t documentor:latest ."
            }
        }
        stage('Deploy To Artifactory') {
            steps {
                sh "printenv"
                sh "deploying somewhere ...."
            }
        }
    }

    post {
        failure {
            sh "docker rmi documentor:latest"
        }
    }
}