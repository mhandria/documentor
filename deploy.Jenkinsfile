pipeline {
    agent {label 'docker'}

    parameters {
        gitParameter name: 'TAG', type: 'PT_TAG', defaultValue: 'master'
    }
    stages {
        stage('Get Tag') {
            steps {
                checkout([$class: 'GitSCM',
                          branches: [[name: "${params.TAG}"]],
                          doGenerateSubmoduleConfigurations: false,
                          extensions: [],
                          gitTool: 'Default',
                          submoduleCfg: [],
                          userRemoteConfigs: [[url: 'https://github.com/jenkinsci/git-parameter-plugin.git']]
                        ])
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
            when {
                expression { params.TAG != 'master' }
            }
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