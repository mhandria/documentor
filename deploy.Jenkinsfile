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
                echo 'Deploying... '
            }
        }
    }
}