from diagrams import Cluster, Diagram
from diagrams.aws.compute import ECS, EKS, Lambda, Fargate

from diagrams.aws.database import Redshift
from diagrams.aws.integration import Eventbridge
from diagrams.aws.storage import S3
from diagrams.aws.management import Cloudwatch


with Diagram("CloudQuery On ECS", show=False):
    source = Eventbridge("Schedule")

    with Cluster("ECS Cluster"):
        with Cluster("ECS Tasks"):
            workers = [Fargate("Fargate"),]

    store = S3("Destination Bucket")
    cwl = Cloudwatch("Structured logs\n from sync")

    source >> workers >> store
    workers >> cwl
