pipeline {
    agent any

    stages {
        stage('Clone Code') {
            steps {
                // Clone your repo
                git branch: 'main', url: 'https://github.com/Rk-112233/gitpractice.git'

            }
        }

        stage('Run Hello World') {
            steps {
                // Run the Go file
                sh 'go run test.go' 
            }
        }
    }
}
