output "bucket_name" {
  description = "Name of the created S3 bucket"
  value       = aws_s3_bucket.app_bucket.bucket
}

output "cloudfront_url" {
  description = "CloudFront distribution URL (if enabled)"
  value       = var.enable_cloudfront ? aws_cloudfront_distribution.s3_distribution[0].domain_name : null
}
