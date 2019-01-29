pipeline {
    agent {label 'docker'}

    stages {
        stage('Get Tag') {
            steps {
                echo "specifying the tags here"
                sh "git describe --abbrev=0 --tags"
            }
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