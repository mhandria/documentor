pipeline {
    agent {label 'docker'}

    stages {
        stage('Server Initial Clean Up') {
            steps {
                sh "docker system prune -f"
            }
        }
        stage('Get Tag') {
            sh "git describe --abbrev=0 --tags"
        }
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
        stage('After Build Clean Up') {
            steps {
                sh "docker rmi documentor:latest"
                sh "docker system prune -f"
            }
        }
    }
}