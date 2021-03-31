# cloud_net
Cloud Netoworking


## Usage:
Create a Service Account and download the key:
https://cloud.google.com/docs/authentication/production#create_service_account
Ensure that the service account has the rights to View Networking Resources.
Roles: NetworkAccountViewer should be enough

Pass Credential via environment variables(use real path and key):
```bash
export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/my-key.json"
```
go run main.go -project_id xxxxxxx
