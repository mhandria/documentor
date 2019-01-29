pipeline {
    agent {label 'docker'}

    stages {
        stage('Get Tag') {
            steps {
                echo "Specifying the tags here"
                sh "git describe --abbrev=0 --tags"
            }
        }
        stage('Build') {
            steps {
                echo "building"
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

    post {
        always {
            deleteDir()
        }
        failure {
            sh "docker system prune -f"
        }
    }
}