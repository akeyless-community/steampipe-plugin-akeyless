connection "akeyless" {
  plugin = "akeyless"  

  # Defines the type of access. Supported types: "api_key", "aws_iam", "azure_ad", "gcp", "universal_identity", "k8s", "jwt".
  #access_type = ""  

  # The access ID for authentication, required for all access types.
  #access_id = ""  

  # The access key or secret, paired with access_id.
  #access_key = ""  

  # API URL for the Akeyless Gateway, the default URL is https://api.akeyless.io.
  #api_url = ""  

  # JSON Web Token for JWT-based authentication.
  #jwt = ""  

  # User identity token for Universal Identity™ authentication.
  #uid_token = ""  

  # Audience for GCP authentication.
  #gcp_audience = ""  

  # Object ID for Azure AD authentication.
  #azure_object_id = ""  

  # Kubernetes service account token for Kubernetes-based authentication.
  #k8s_service_account_token = ""  

  # Name of the Kubernetes auth config.
  #k8s_auth_config_name = ""  

  # CA certificate for TLS verification of the Akeyless Gateway.
  #gateway_ca_cert = ""  
}
