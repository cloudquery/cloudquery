resource "aws_eks_cluster" "aws_eks_clusters_cluster" {
  name     = "eks-${var.test_prefix}${var.test_suffix}"
  role_arn = aws_iam_role.aws_eks_clusters_iam_role.arn

  vpc_config {
    subnet_ids = [
      aws_subnet.aws_vpc_subnet2.id,
    aws_subnet.aws_vpc_subnet3.id]
  }

  depends_on = [
    aws_iam_role_policy_attachment.aws_eks_clusters_cluster_policy,
    aws_iam_role_policy_attachment.aws_eks_clusters_AmazonEKSVPCResourceController,
  ]
}

output "endpoint" {
  value = aws_eks_cluster.aws_eks_clusters_cluster.endpoint
}

resource "aws_iam_role" "aws_eks_clusters_iam_role" {
  name = "eks-role-${var.test_prefix}${var.test_suffix}"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "eks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "aws_eks_clusters_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.aws_eks_clusters_iam_role.name
}

resource "aws_iam_role_policy_attachment" "aws_eks_clusters_AmazonEKSVPCResourceController" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController"
  role       = aws_iam_role.aws_eks_clusters_iam_role.name
}