pipeline {
    agent {label 'docker'}

    stages {
        
        stage('Build') {
            steps {
                sh "docker build -t documentor:build ."
            }
        }

        stage('Deploy To Artifactory') {
            when {
                expression { env.CHANGE_TARGET == null }
            }
            steps {
                sh "docker tag documentor:build documentor:${env.TAG_NAME}"
                echo "deploying somewhere ...."
            }
        }
    }

    post {
        always {
            deleteDir()
        }
        failure {
            sh "docker rmi documentor:latest"
        }
    }
}