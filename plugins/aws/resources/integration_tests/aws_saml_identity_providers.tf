resource "aws_iam_saml_provider" "default_saml" {
  name = "${var.test_prefix}${var.test_suffix}"
  saml_metadata_document = <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<EntityDescriptor entityID="https://client.mydomain.com:443/webconsole" xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata">
<SPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol" WantAssertionsSigned="true">
<AssertionConsumerService isDefault="true" index="0" Location="https://client.mydomain.com:443/webconsole/samlAcsIdpInitCallback.do?samlAppKey=NjZEOUQ1RDRCQjE1NEI0" Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"/>
<md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location=" https://client.mydomain.com:443/webconsole/server/SAMLSingleLogout?samlAppKey=MzU2MkNDQTFBQzczNEZG" ResponseLocation="https://client.mydomain.com:443/webconsole/server/SAMLSingleLogout"/>
<md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location=" https://client.mydomain.com:443/webconsole/server/SAMLSingleLogout?samlAppKey=MzU2MkNDQTFBQzczNEZG"ResponseLocation="https://client.mydomain.com:443/webconsole/server/SAMLSingleLogout"/>
<KeyDescriptor>
<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
<ds:X509Data>
<ds:X509Certificate>encoded_certificate</ds:X509Certificate>
</ds:X509Data>
</ds:KeyInfo>
</KeyDescriptor>
<NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:entity</NameIDFormat>
</SPSSODescriptor>
</EntityDescriptor>
EOF
}