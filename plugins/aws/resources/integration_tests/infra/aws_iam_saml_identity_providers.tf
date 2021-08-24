resource "aws_iam_saml_provider" "default_saml" {
  name = "saml${var.test_prefix}${var.test_suffix}"
  saml_metadata_document = <<EOF
<md:EntityDescriptor entityID="http://www.example.com" ID="_58eb6efc-1f19-431b-8146-3fef71f908d0"
                     xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata">
    <Signature xmlns="http://www.w3.org/2000/09/xmldsig#">
        <SignedInfo>
            <CanonicalizationMethod Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/>
            <SignatureMethod Algorithm="http://www.w3.org/2000/09/xmldsig#rsa-sha1"/>
            <Reference URI="#_58eb6efc-1f19-431b-8146-3fef71f908d0">
                <Transforms>
                    <Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/>
                    <Transform Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#">
                        <InclusiveNamespaces PrefixList="#default md saml ds xs xsi"
                                             xmlns="http://www.w3.org/2001/10/xml-exc-c14n#"/>
                    </Transform>
                </Transforms>
                <DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1"/>
                <DigestValue>s0NN/wBccBPbajTQtI+Eyf88kPg=</DigestValue>
            </Reference>
        </SignedInfo>
        <SignatureValue>
            Vp1p+qgJ+GSZ1STdig8mhgcE41sI2//zHEjlDVX5PxTZLjdOwRkN6KcNA0fJNAVABMYC4Hhaw22iBwgtAzD2qmiuiBMBemJFIdPbbJ07zVQ1BZpA2Bpfx7Ymjai9ENwn3m6iGhNMRxa5NlntQCNH6ua2KAZxflX/3jUlQjt63w89hake8iIYVBtRVe6LIH52Ic9lzsmdn/cs6jXJ1BWGMkI5z41SCkf+80x3u5DDJ1ljJGYHPFdWXGClho8Wq1iK2PgLOWnaY5bf8gtpdMk2Shn2sfwN2SgQtzFo+YJ0YfXwZcoPx3p1WzYa6V5z0dwYTH24qxf15JlvwJWPmhwToA==
        </SignatureValue>
        <KeyInfo>
            <X509Data>
                <X509Certificate>
                    MIIDATCCAemgAwIBAgIQdPDr/iI1jbhDMTj5VYya+TANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDEwt3d3cuaWRwLmNvbTAeFw0xMzExMjIwODIwNTJaFw00OTEyMzExNDAwMDBaMBYxFDASBgNVBAMTC3d3dy5pZHAuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAi0XJRLDrcbSyqUd8XG4BgxObQMYLAkENlmJOsAEpl1xMabUiq1X4v0Fc8ZaCpUE3fFGENMEWgBjnQUUE0WtVUh5JPMsukolf9qljbJkCkvHXH3O4Uen7vA2oNQWt4bK96SpXADpZKFvpk4D7btKOgU/NamjiqwHI4fI8kFJKwKBJchRPUQdC4ljRRmGIrSnpY+t25/d3KGXwbe9Z2MGGy2hyA0tgOWuchIK+1vAKKBUh9nDEXfr80+xW680w5TqHyDcqbWvQsXXhH0yZLfINKNS6/IojHPsBy7tf36Ck9H5Pw+1PPu6NzBFSz5ZkC8KzrS6vuZXc/ImYrnheMQsqqQIDAQABo0swSTBHBgNVHQEEQDA+gBD4dY4MCPEmG4sxZrcni8vtoRgwFjEUMBIGA1UEAxMLd3d3LmlkcC5jb22CEHTw6/4iNY24QzE4+VWMmvkwDQYJKoZIhvcNAQELBQADggEBABhak2aR84MCdyXO4AKOQvZybsCMdhRq2i1i0WhD4/xe7Ry5haC6TeXIp8Q4cC3MzsrDal74xHI714BW0loafpHAsXfd9EvkKTVaJ+1Zpe16+SsTL4upS1cGydigqwUzsdpGck4wI1moJ9477O+46If2gF27u9Cdk7Onxe/5dwLIxWmkVRdbQIH5GsKUeAjOdRQmy+X1MX6KyRoaCwWGYwxi5Sa+r+3AtDvD4BX0EJGKFZeeM3J/yMpYh/75aN0cFQfDEdJ7C5NE0vonidE0QtIFvsoWtZUtur2fiW7yBxse38TPQsi2r6A6c/TZsZ5bq31yh3gr3kSN62H8iVKLQLA=
                </X509Certificate>
            </X509Data>
        </KeyInfo>
    </Signature>
    <md:IDPSSODescriptor ID="_04494e78-cc4f-404b-9fa8-16373b073228"
                         protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"
                         WantAuthnRequestsSigned="true">
        <md:KeyDescriptor>
            <KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#">
                <X509Data>
                    <X509Certificate>
                        MIIDATCCAemgAwIBAgIQdPDr/iI1jbhDMTj5VYya+TANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDEwt3d3cuaWRwLmNvbTAeFw0xMzExMjIwODIwNTJaFw00OTEyMzExNDAwMDBaMBYxFDASBgNVBAMTC3d3dy5pZHAuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAi0XJRLDrcbSyqUd8XG4BgxObQMYLAkENlmJOsAEpl1xMabUiq1X4v0Fc8ZaCpUE3fFGENMEWgBjnQUUE0WtVUh5JPMsukolf9qljbJkCkvHXH3O4Uen7vA2oNQWt4bK96SpXADpZKFvpk4D7btKOgU/NamjiqwHI4fI8kFJKwKBJchRPUQdC4ljRRmGIrSnpY+t25/d3KGXwbe9Z2MGGy2hyA0tgOWuchIK+1vAKKBUh9nDEXfr80+xW680w5TqHyDcqbWvQsXXhH0yZLfINKNS6/IojHPsBy7tf36Ck9H5Pw+1PPu6NzBFSz5ZkC8KzrS6vuZXc/ImYrnheMQsqqQIDAQABo0swSTBHBgNVHQEEQDA+gBD4dY4MCPEmG4sxZrcni8vtoRgwFjEUMBIGA1UEAxMLd3d3LmlkcC5jb22CEHTw6/4iNY24QzE4+VWMmvkwDQYJKoZIhvcNAQELBQADggEBABhak2aR84MCdyXO4AKOQvZybsCMdhRq2i1i0WhD4/xe7Ry5haC6TeXIp8Q4cC3MzsrDal74xHI714BW0loafpHAsXfd9EvkKTVaJ+1Zpe16+SsTL4upS1cGydigqwUzsdpGck4wI1moJ9477O+46If2gF27u9Cdk7Onxe/5dwLIxWmkVRdbQIH5GsKUeAjOdRQmy+X1MX6KyRoaCwWGYwxi5Sa+r+3AtDvD4BX0EJGKFZeeM3J/yMpYh/75aN0cFQfDEdJ7C5NE0vonidE0QtIFvsoWtZUtur2fiW7yBxse38TPQsi2r6A6c/TZsZ5bq31yh3gr3kSN62H8iVKLQLA=
                    </X509Certificate>
                </X509Data>
            </KeyInfo>
        </md:KeyDescriptor>
        <md:ArtifactResolutionService Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"
                                      Location="https://www.example.com/ArtifactResolutionService" index="1"
                                      isDefault="true"/>
        <md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:transient</md:NameIDFormat>
        <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
                                Location="https://www.example.com/SSOService"/>
    </md:IDPSSODescriptor>
    <md:Organization>
        <md:OrganizationName xml:lang="en">example</md:OrganizationName>
        <md:OrganizationDisplayName xml:lang="en">example</md:OrganizationDisplayName>
        <md:OrganizationURL xml:lang="en">www.example.com</md:OrganizationURL>
    </md:Organization>
    <md:ContactPerson contactType="technical">
        <md:GivenName>Test</md:GivenName>
        <md:SurName>Test</md:SurName>
        <md:EmailAddress>test.test@example.com</md:EmailAddress>
    </md:ContactPerson>
</md:EntityDescriptor>
EOF
}