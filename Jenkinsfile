pipeline {
    agent any
    tools {
        go 'go-1.15'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'cd cidr_convert_api/go/'
                sh 'pwd'
                sh 'ls'
                sh 'go build'
            }
        }
        stage('Test') {
            /*environment {
                CODECOV_TOKEN = credentials('codecov_token')
            }*/
            steps {
                sh 'go version'
                sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.15.6'
                sh 'golangci-lint run'
            }
        }
        stage('Code Analysis') {
            steps {
                 sh 'golangci-lint run'
            }
        }
        stage('Release') {
            when {
                buildingTag()
            }
            environment {
                GITHUB_TOKEN = credentials('github_token')
            }
            steps {
                sh 'curl -sL https://git.io/goreleaser | bash'
            }
        }
    }
}
