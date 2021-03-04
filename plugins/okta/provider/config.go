package provider

const configYaml = `
  - name: okta
    domain: https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com
    resources:
      - name: users
      - name: applications
	  - name: userTypes
      - name: groups`
