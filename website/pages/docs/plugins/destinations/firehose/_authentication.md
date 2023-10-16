Authentication is handled by the AWS SDK. Credentials and configurations are sourced from the environment. Credentials are sourced in the following order:

1. Environment variables.
Static Credentials (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`)
Web Identity Token (`AWS_WEB_IDENTITY_TOKEN_FILE`)
2. Shared configuration files.
SDK defaults to `credentials` file under `.aws` folder that is placed in the home folder on your computer.
SDK defaults to `config` file under `.aws` folder that is placed in the home folder on your computer.
3. If your application uses an ECS task definition or RunTask API operation, IAM role for tasks.
4. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
