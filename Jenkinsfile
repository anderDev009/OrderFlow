pipeline {
    agent any
    environment {
        IMAGE_NAME = "richarddev009/orderflow"
        IMAGE_TAG = "${BUILD_NUMBER}"
        REGISTRY = "docker.io"
    }
    stages {

        stage("Checkout"){
            steps{
                checkout scm
            }
        }
        stage('Building docker image') {
            steps {
                sh """
                    docker build -t $IMAGE_NAME:$IMAGE_TAG .
                    docker tag $IMAGE_NAME:$IMAGE_TAG $image_name:latest 
                """
            }
        }
        stage("Logging dockerhub"){
            steps{
                sh """
                    echo $PERSONAL_DOCKER_PASSW | docker login -u $PERSONAL_DOCKER_USER --password-stdin
                """
            }
        }
        stage("Pushing image"){
            steps{
                sh """
                docker push $IMAGE_NAME:$IMAGE_TAG
                docker push $IMAGE_NAME:latest
                """
            }
        }
        // stage("Deploy proxmox VM"){
        //     steps{

        //     }
        // }
    }
    post{
        success{
            echo "Despliegue exitoso"
        }
        failure{
            echo "Error en el pipeline"
        }
    }
}