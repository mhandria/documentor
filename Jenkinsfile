pipeline {
    agent {label 'docker'}

    stages {
        stage('Server Initial Clean Up') {
            steps {
                sh "docker system prune -f"
            }
        }
        stage('Build') {
            steps {
                sh "docker build -t documentor:latest ."
            }
        }
        stage('After Build Clean Up') {
            steps {
                sh "docker rmi documentor:latest"
                sh "docker system prune -f"
            }
        }
    }
}