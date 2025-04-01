variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "bucket_prefix" {
  description = "Prefix for S3 bucket name"
  type        = string
  default     = "my-app"
}

variable "enable_cloudfront" {
  description = "Whether to enable CloudFront distribution"
  type        = bool
  default     = false
}
