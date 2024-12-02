package client_test

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/snowflake/v4/client"
	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpec_DSN(t *testing.T) {
	tests := []struct {
		Spec    client.Spec
		Extra   map[string]*string
		WantDSN string
		WantErr string
	}{
		{
			Spec: client.Spec{
				ConnectionString: "user@acc/db/schema?warehouse=wh",
				PrivateKey:       "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDAWKgfyKnlEMAm\nbR5Kc5V//3LFJp0rQyR5gVRIxvsU/9RaWi+c/9FBCPVGpi4yJ5TGRRKxoZwX/yCg\nnE76Vxatz26DwM3qagAfNDjtKlyqdLSrMWUNTSj8WVKPd9sPnDErYvIw21Nnx4Sj\nEt0Zku1C+K+GVw0J83Fu7d0dVB3XM1KKguY8+POQ5gs+kOef9p4yaFkwqqCeigdb\nIU7XSw/6A5jFWajKzpdL1mH8H2iidMSVWJ8Pngx4mdk5dVbrLAhgyuIfgsOSdane\nX851+7qozlog67JR6wLbIGWPO7RpJMCYkR0Z2n5BHmpnqJ95qFTOWIAH24XKdeib\n+H2X8+DFAgMBAAECggEAL53zFxU97AGS1CB42n0VQl/6qXB3AcYIMlVQSIkMEQWJ\nbEm91kvlYYiGchRDRPLUC6Z/a36i7jTgfqpbifGD4YkD5rWVNIZD2/W5bwso8CDe\nti/PALU8o4Y4YGCPWGS2LnO7Gdm+Iue7gAR0PHfJaWY3y9XimjdMemYD8pYHoiW6\nATQ4XkLskvlqiLnyNMdb1ByTGxgB8778O/HsoMyNdzYeWgxyZtfaa/CBr0xQr92Q\n9CfcB0YizK0I+nygP4RR5CFPEM+zQ9VEeFHTh4yci0A4+VYs7NRpSUrUW7TVeu4d\nAGKHnooe/EL4DtdcNZVUipZLLy6lHCEAuWvWVAVQEQKBgQDv/Qzj3tKuXOENBmNE\nLluS5dIE+KYMa9chioctiJqIWT9rrHz2/EU7b1W5mxby7WJFppwMzFHJP6GgnFaw\niYgK54YMsPv/5WrXftN4vC3vuw1pOUZIPT41e8Uu30lFcl/crpEfwA3Rg/2m57Fe\nptPPKn0+y50JiQIo3NxcY6LsNQKBgQDNLeONgojsGwen1PoaEtEvAdYhmwI3HVbV\nFyKNifO7T7eISM4rwMP3FstykHzpcTwno/WrC5+jCeNc7q/U1AEudLu2YYpJKRlZ\nbEv4rAfZeJp08f6lxKOuI5SLL7iZ3MqsoAAnv1kwYCztlcYDt4c3LVYTub5tcWHv\n34rWi4cUUQKBgC9sy1pQk0O/uP2Q8Jbtrk0GO42d8Xps6TOIo5P89cTSFjVZ/cv1\nKF1JcCBgpJVXEd9/wEDLM7JYb8FEg+EZHJhDDnt9kh8MoCN7vaCTV2STi1/q4Jev\n+pYpIltT5q/hnU4H9UfX9SMdOUf9a1CwGRVMaTm6lQroV1Pp6WYcjnqtAoGARUO0\nidUDPBFz6ChxtdOcYm4QR4/4k3qIEa+ZroZfjWA/6PYLA6IzhXpge/Bi+ruLPyaO\njIuD/Jod8wVwvjxDmdc2dz8+W6xQLmvsyanpjHS2T7xR5swXJXZFcydM/kQW92ec\nJc7m4PnWsO3axu5x6yKW6FnP+0pHcZ7ZU8wOccECgYA2rkrQVk9qHqBzymZFj47I\nQwD65jlzNeK9xMeKHivsJc4K9cmwec4jRcsu8gEYcWIbhepjHUA5M/4mO0++bnsP\nPP2xIiqseAOt4rI0wTeXqY8AM6YvbfdMq+tN+BlXix0WBEGNFGpab3B9bmv8UK0L\njhuCua5QHoIX6/fasAJ88w==\n-----END PRIVATE KEY-----\n",
			},
			WantDSN: "user@acc/db/schema?warehouse=wh&authenticator=snowflake_jwt&privateKey=MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDAWKgfyKnlEMAmbR5Kc5V__3LFJp0rQyR5gVRIxvsU_9RaWi-c_9FBCPVGpi4yJ5TGRRKxoZwX_yCgnE76Vxatz26DwM3qagAfNDjtKlyqdLSrMWUNTSj8WVKPd9sPnDErYvIw21Nnx4SjEt0Zku1C-K-GVw0J83Fu7d0dVB3XM1KKguY8-POQ5gs-kOef9p4yaFkwqqCeigdbIU7XSw_6A5jFWajKzpdL1mH8H2iidMSVWJ8Pngx4mdk5dVbrLAhgyuIfgsOSdaneX851-7qozlog67JR6wLbIGWPO7RpJMCYkR0Z2n5BHmpnqJ95qFTOWIAH24XKdeib-H2X8-DFAgMBAAECggEAL53zFxU97AGS1CB42n0VQl_6qXB3AcYIMlVQSIkMEQWJbEm91kvlYYiGchRDRPLUC6Z_a36i7jTgfqpbifGD4YkD5rWVNIZD2_W5bwso8CDeti_PALU8o4Y4YGCPWGS2LnO7Gdm-Iue7gAR0PHfJaWY3y9XimjdMemYD8pYHoiW6ATQ4XkLskvlqiLnyNMdb1ByTGxgB8778O_HsoMyNdzYeWgxyZtfaa_CBr0xQr92Q9CfcB0YizK0I-nygP4RR5CFPEM-zQ9VEeFHTh4yci0A4-VYs7NRpSUrUW7TVeu4dAGKHnooe_EL4DtdcNZVUipZLLy6lHCEAuWvWVAVQEQKBgQDv_Qzj3tKuXOENBmNELluS5dIE-KYMa9chioctiJqIWT9rrHz2_EU7b1W5mxby7WJFppwMzFHJP6GgnFawiYgK54YMsPv_5WrXftN4vC3vuw1pOUZIPT41e8Uu30lFcl_crpEfwA3Rg_2m57FeptPPKn0-y50JiQIo3NxcY6LsNQKBgQDNLeONgojsGwen1PoaEtEvAdYhmwI3HVbVFyKNifO7T7eISM4rwMP3FstykHzpcTwno_WrC5-jCeNc7q_U1AEudLu2YYpJKRlZbEv4rAfZeJp08f6lxKOuI5SLL7iZ3MqsoAAnv1kwYCztlcYDt4c3LVYTub5tcWHv34rWi4cUUQKBgC9sy1pQk0O_uP2Q8Jbtrk0GO42d8Xps6TOIo5P89cTSFjVZ_cv1KF1JcCBgpJVXEd9_wEDLM7JYb8FEg-EZHJhDDnt9kh8MoCN7vaCTV2STi1_q4Jev-pYpIltT5q_hnU4H9UfX9SMdOUf9a1CwGRVMaTm6lQroV1Pp6WYcjnqtAoGARUO0idUDPBFz6ChxtdOcYm4QR4_4k3qIEa-ZroZfjWA_6PYLA6IzhXpge_Bi-ruLPyaOjIuD_Jod8wVwvjxDmdc2dz8-W6xQLmvsyanpjHS2T7xR5swXJXZFcydM_kQW92ecJc7m4PnWsO3axu5x6yKW6FnP-0pHcZ7ZU8wOccECgYA2rkrQVk9qHqBzymZFj47IQwD65jlzNeK9xMeKHivsJc4K9cmwec4jRcsu8gEYcWIbhepjHUA5M_4mO0--bnsPPP2xIiqseAOt4rI0wTeXqY8AM6YvbfdMq-tN-BlXix0WBEGNFGpab3B9bmv8UK0LjhuCua5QHoIX6_fasAJ88w==",
		},
		{
			Spec: client.Spec{
				ConnectionString: "user@acc/db/schema?warehouse=wh",
				// This is the key above but with all line breaks replaced with
				// spaces, which is invalid, but we're going to repair the key.
				PrivateKey: "-----BEGIN PRIVATE KEY-----  MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDAWKgfyKnlEMAm  bR5Kc5V//3LFJp0rQyR5gVRIxvsU/9RaWi+c/9FBCPVGpi4yJ5TGRRKxoZwX/yCg  nE76Vxatz26DwM3qagAfNDjtKlyqdLSrMWUNTSj8WVKPd9sPnDErYvIw21Nnx4Sj  Et0Zku1C+K+GVw0J83Fu7d0dVB3XM1KKguY8+POQ5gs+kOef9p4yaFkwqqCeigdb  IU7XSw/6A5jFWajKzpdL1mH8H2iidMSVWJ8Pngx4mdk5dVbrLAhgyuIfgsOSdane  X851+7qozlog67JR6wLbIGWPO7RpJMCYkR0Z2n5BHmpnqJ95qFTOWIAH24XKdeib  +H2X8+DFAgMBAAECggEAL53zFxU97AGS1CB42n0VQl/6qXB3AcYIMlVQSIkMEQWJ  bEm91kvlYYiGchRDRPLUC6Z/a36i7jTgfqpbifGD4YkD5rWVNIZD2/W5bwso8CDe  ti/PALU8o4Y4YGCPWGS2LnO7Gdm+Iue7gAR0PHfJaWY3y9XimjdMemYD8pYHoiW6  ATQ4XkLskvlqiLnyNMdb1ByTGxgB8778O/HsoMyNdzYeWgxyZtfaa/CBr0xQr92Q  9CfcB0YizK0I+nygP4RR5CFPEM+zQ9VEeFHTh4yci0A4+VYs7NRpSUrUW7TVeu4d  AGKHnooe/EL4DtdcNZVUipZLLy6lHCEAuWvWVAVQEQKBgQDv/Qzj3tKuXOENBmNE  LluS5dIE+KYMa9chioctiJqIWT9rrHz2/EU7b1W5mxby7WJFppwMzFHJP6GgnFaw  iYgK54YMsPv/5WrXftN4vC3vuw1pOUZIPT41e8Uu30lFcl/crpEfwA3Rg/2m57Fe  ptPPKn0+y50JiQIo3NxcY6LsNQKBgQDNLeONgojsGwen1PoaEtEvAdYhmwI3HVbV  FyKNifO7T7eISM4rwMP3FstykHzpcTwno/WrC5+jCeNc7q/U1AEudLu2YYpJKRlZ  bEv4rAfZeJp08f6lxKOuI5SLL7iZ3MqsoAAnv1kwYCztlcYDt4c3LVYTub5tcWHv  34rWi4cUUQKBgC9sy1pQk0O/uP2Q8Jbtrk0GO42d8Xps6TOIo5P89cTSFjVZ/cv1  KF1JcCBgpJVXEd9/wEDLM7JYb8FEg+EZHJhDDnt9kh8MoCN7vaCTV2STi1/q4Jev  +pYpIltT5q/hnU4H9UfX9SMdOUf9a1CwGRVMaTm6lQroV1Pp6WYcjnqtAoGARUO0  idUDPBFz6ChxtdOcYm4QR4/4k3qIEa+ZroZfjWA/6PYLA6IzhXpge/Bi+ruLPyaO  jIuD/Jod8wVwvjxDmdc2dz8+W6xQLmvsyanpjHS2T7xR5swXJXZFcydM/kQW92ec  Jc7m4PnWsO3axu5x6yKW6FnP+0pHcZ7ZU8wOccECgYA2rkrQVk9qHqBzymZFj47I  QwD65jlzNeK9xMeKHivsJc4K9cmwec4jRcsu8gEYcWIbhepjHUA5M/4mO0++bnsP  PP2xIiqseAOt4rI0wTeXqY8AM6YvbfdMq+tN+BlXix0WBEGNFGpab3B9bmv8UK0L  jhuCua5QHoIX6/fasAJ88w==  -----END PRIVATE KEY-----  ",
			},
			WantDSN: "user@acc/db/schema?warehouse=wh&authenticator=snowflake_jwt&privateKey=MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDAWKgfyKnlEMAmbR5Kc5V__3LFJp0rQyR5gVRIxvsU_9RaWi-c_9FBCPVGpi4yJ5TGRRKxoZwX_yCgnE76Vxatz26DwM3qagAfNDjtKlyqdLSrMWUNTSj8WVKPd9sPnDErYvIw21Nnx4SjEt0Zku1C-K-GVw0J83Fu7d0dVB3XM1KKguY8-POQ5gs-kOef9p4yaFkwqqCeigdbIU7XSw_6A5jFWajKzpdL1mH8H2iidMSVWJ8Pngx4mdk5dVbrLAhgyuIfgsOSdaneX851-7qozlog67JR6wLbIGWPO7RpJMCYkR0Z2n5BHmpnqJ95qFTOWIAH24XKdeib-H2X8-DFAgMBAAECggEAL53zFxU97AGS1CB42n0VQl_6qXB3AcYIMlVQSIkMEQWJbEm91kvlYYiGchRDRPLUC6Z_a36i7jTgfqpbifGD4YkD5rWVNIZD2_W5bwso8CDeti_PALU8o4Y4YGCPWGS2LnO7Gdm-Iue7gAR0PHfJaWY3y9XimjdMemYD8pYHoiW6ATQ4XkLskvlqiLnyNMdb1ByTGxgB8778O_HsoMyNdzYeWgxyZtfaa_CBr0xQr92Q9CfcB0YizK0I-nygP4RR5CFPEM-zQ9VEeFHTh4yci0A4-VYs7NRpSUrUW7TVeu4dAGKHnooe_EL4DtdcNZVUipZLLy6lHCEAuWvWVAVQEQKBgQDv_Qzj3tKuXOENBmNELluS5dIE-KYMa9chioctiJqIWT9rrHz2_EU7b1W5mxby7WJFppwMzFHJP6GgnFawiYgK54YMsPv_5WrXftN4vC3vuw1pOUZIPT41e8Uu30lFcl_crpEfwA3Rg_2m57FeptPPKn0-y50JiQIo3NxcY6LsNQKBgQDNLeONgojsGwen1PoaEtEvAdYhmwI3HVbVFyKNifO7T7eISM4rwMP3FstykHzpcTwno_WrC5-jCeNc7q_U1AEudLu2YYpJKRlZbEv4rAfZeJp08f6lxKOuI5SLL7iZ3MqsoAAnv1kwYCztlcYDt4c3LVYTub5tcWHv34rWi4cUUQKBgC9sy1pQk0O_uP2Q8Jbtrk0GO42d8Xps6TOIo5P89cTSFjVZ_cv1KF1JcCBgpJVXEd9_wEDLM7JYb8FEg-EZHJhDDnt9kh8MoCN7vaCTV2STi1_q4Jev-pYpIltT5q_hnU4H9UfX9SMdOUf9a1CwGRVMaTm6lQroV1Pp6WYcjnqtAoGARUO0idUDPBFz6ChxtdOcYm4QR4_4k3qIEa-ZroZfjWA_6PYLA6IzhXpge_Bi-ruLPyaOjIuD_Jod8wVwvjxDmdc2dz8-W6xQLmvsyanpjHS2T7xR5swXJXZFcydM_kQW92ecJc7m4PnWsO3axu5x6yKW6FnP-0pHcZ7ZU8wOccECgYA2rkrQVk9qHqBzymZFj47IQwD65jlzNeK9xMeKHivsJc4K9cmwec4jRcsu8gEYcWIbhepjHUA5M_4mO0--bnsPPP2xIiqseAOt4rI0wTeXqY8AM6YvbfdMq-tN-BlXix0WBEGNFGpab3B9bmv8UK0LjhuCua5QHoIX6_fasAJ88w==",
		},
		{
			Spec: client.Spec{
				ConnectionString: "user@acc/db/schema?warehouse=wh",
				PrivateKey:       "-----BEGIN ENCRYPTED PRIVATE KEY-----  MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDAWKgfyKnlEMAm  bR5Kc5V//3LFJp0rQyR5gVRIxvsU/9RaWi+c/9FBCPVGpi4yJ5TGRRKxoZwX/yCg  nE76Vxatz26DwM3qagAfNDjtKlyqdLSrMWUNTSj8WVKPd9sPnDErYvIw21Nnx4Sj  Et0Zku1C+K+GVw0J83Fu7d0dVB3XM1KKguY8+POQ5gs+kOef9p4yaFkwqqCeigdb  IU7XSw/6A5jFWajKzpdL1mH8H2iidMSVWJ8Pngx4mdk5dVbrLAhgyuIfgsOSdane  X851+7qozlog67JR6wLbIGWPO7RpJMCYkR0Z2n5BHmpnqJ95qFTOWIAH24XKdeib  +H2X8+DFAgMBAAECggEAL53zFxU97AGS1CB42n0VQl/6qXB3AcYIMlVQSIkMEQWJ  bEm91kvlYYiGchRDRPLUC6Z/a36i7jTgfqpbifGD4YkD5rWVNIZD2/W5bwso8CDe  ti/PALU8o4Y4YGCPWGS2LnO7Gdm+Iue7gAR0PHfJaWY3y9XimjdMemYD8pYHoiW6  ATQ4XkLskvlqiLnyNMdb1ByTGxgB8778O/HsoMyNdzYeWgxyZtfaa/CBr0xQr92Q  9CfcB0YizK0I+nygP4RR5CFPEM+zQ9VEeFHTh4yci0A4+VYs7NRpSUrUW7TVeu4d  AGKHnooe/EL4DtdcNZVUipZLLy6lHCEAuWvWVAVQEQKBgQDv/Qzj3tKuXOENBmNE  LluS5dIE+KYMa9chioctiJqIWT9rrHz2/EU7b1W5mxby7WJFppwMzFHJP6GgnFaw  iYgK54YMsPv/5WrXftN4vC3vuw1pOUZIPT41e8Uu30lFcl/crpEfwA3Rg/2m57Fe  ptPPKn0+y50JiQIo3NxcY6LsNQKBgQDNLeONgojsGwen1PoaEtEvAdYhmwI3HVbV  FyKNifO7T7eISM4rwMP3FstykHzpcTwno/WrC5+jCeNc7q/U1AEudLu2YYpJKRlZ  bEv4rAfZeJp08f6lxKOuI5SLL7iZ3MqsoAAnv1kwYCztlcYDt4c3LVYTub5tcWHv  34rWi4cUUQKBgC9sy1pQk0O/uP2Q8Jbtrk0GO42d8Xps6TOIo5P89cTSFjVZ/cv1  KF1JcCBgpJVXEd9/wEDLM7JYb8FEg+EZHJhDDnt9kh8MoCN7vaCTV2STi1/q4Jev  +pYpIltT5q/hnU4H9UfX9SMdOUf9a1CwGRVMaTm6lQroV1Pp6WYcjnqtAoGARUO0  idUDPBFz6ChxtdOcYm4QR4/4k3qIEa+ZroZfjWA/6PYLA6IzhXpge/Bi+ruLPyaO  jIuD/Jod8wVwvjxDmdc2dz8+W6xQLmvsyanpjHS2T7xR5swXJXZFcydM/kQW92ec  Jc7m4PnWsO3axu5x6yKW6FnP+0pHcZ7ZU8wOccECgYA2rkrQVk9qHqBzymZFj47I  QwD65jlzNeK9xMeKHivsJc4K9cmwec4jRcsu8gEYcWIbhepjHUA5M/4mO0++bnsP  PP2xIiqseAOt4rI0wTeXqY8AM6YvbfdMq+tN+BlXix0WBEGNFGpab3B9bmv8UK0L  jhuCua5QHoIX6/fasAJ88w==  -----END PRIVATE KEY-----  ",
			},
			WantErr: "private_key: encrypted private keys are not supported, use decrypted private key",
		},
		{
			Spec: client.Spec{
				ConnectionString: "user@acc/db/schema?warehouse=wh",
				PrivateKey:       "-----BEGIN WEIRD PRIVATE KEY-----  MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDAWKgfyKnlEMAm  bR5Kc5V//3LFJp0rQyR5gVRIxvsU/9RaWi+c/9FBCPVGpi4yJ5TGRRKxoZwX/yCg  nE76Vxatz26DwM3qagAfNDjtKlyqdLSrMWUNTSj8WVKPd9sPnDErYvIw21Nnx4Sj  Et0Zku1C+K+GVw0J83Fu7d0dVB3XM1KKguY8+POQ5gs+kOef9p4yaFkwqqCeigdb  IU7XSw/6A5jFWajKzpdL1mH8H2iidMSVWJ8Pngx4mdk5dVbrLAhgyuIfgsOSdane  X851+7qozlog67JR6wLbIGWPO7RpJMCYkR0Z2n5BHmpnqJ95qFTOWIAH24XKdeib  +H2X8+DFAgMBAAECggEAL53zFxU97AGS1CB42n0VQl/6qXB3AcYIMlVQSIkMEQWJ  bEm91kvlYYiGchRDRPLUC6Z/a36i7jTgfqpbifGD4YkD5rWVNIZD2/W5bwso8CDe  ti/PALU8o4Y4YGCPWGS2LnO7Gdm+Iue7gAR0PHfJaWY3y9XimjdMemYD8pYHoiW6  ATQ4XkLskvlqiLnyNMdb1ByTGxgB8778O/HsoMyNdzYeWgxyZtfaa/CBr0xQr92Q  9CfcB0YizK0I+nygP4RR5CFPEM+zQ9VEeFHTh4yci0A4+VYs7NRpSUrUW7TVeu4d  AGKHnooe/EL4DtdcNZVUipZLLy6lHCEAuWvWVAVQEQKBgQDv/Qzj3tKuXOENBmNE  LluS5dIE+KYMa9chioctiJqIWT9rrHz2/EU7b1W5mxby7WJFppwMzFHJP6GgnFaw  iYgK54YMsPv/5WrXftN4vC3vuw1pOUZIPT41e8Uu30lFcl/crpEfwA3Rg/2m57Fe  ptPPKn0+y50JiQIo3NxcY6LsNQKBgQDNLeONgojsGwen1PoaEtEvAdYhmwI3HVbV  FyKNifO7T7eISM4rwMP3FstykHzpcTwno/WrC5+jCeNc7q/U1AEudLu2YYpJKRlZ  bEv4rAfZeJp08f6lxKOuI5SLL7iZ3MqsoAAnv1kwYCztlcYDt4c3LVYTub5tcWHv  34rWi4cUUQKBgC9sy1pQk0O/uP2Q8Jbtrk0GO42d8Xps6TOIo5P89cTSFjVZ/cv1  KF1JcCBgpJVXEd9/wEDLM7JYb8FEg+EZHJhDDnt9kh8MoCN7vaCTV2STi1/q4Jev  +pYpIltT5q/hnU4H9UfX9SMdOUf9a1CwGRVMaTm6lQroV1Pp6WYcjnqtAoGARUO0  idUDPBFz6ChxtdOcYm4QR4/4k3qIEa+ZroZfjWA/6PYLA6IzhXpge/Bi+ruLPyaO  jIuD/Jod8wVwvjxDmdc2dz8+W6xQLmvsyanpjHS2T7xR5swXJXZFcydM/kQW92ec  Jc7m4PnWsO3axu5x6yKW6FnP+0pHcZ7ZU8wOccECgYA2rkrQVk9qHqBzymZFj47I  QwD65jlzNeK9xMeKHivsJc4K9cmwec4jRcsu8gEYcWIbhepjHUA5M/4mO0++bnsP  PP2xIiqseAOt4rI0wTeXqY8AM6YvbfdMq+tN+BlXix0WBEGNFGpab3B9bmv8UK0L  jhuCua5QHoIX6/fasAJ88w==  -----END PRIVATE KEY-----  ",
			},
			WantErr: "private_key: unrecognised start block -----BEGIN WEIRD PRIVATE KEY-----, expected -----BEGIN PRIVATE KEY-----",
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			specjson, err := json.Marshal(test.Spec)
			if err != nil {
				t.Fatalf("marshalling spec %#v: %s", test.Spec, err)
			}

			dsn, err := test.Spec.DSN()
			if err != nil {
				if test.WantErr == "" {
					t.Fatalf("Unwanted error %q from spec: %s", err, specjson)
				} else if !strings.Contains(err.Error(), test.WantErr) {
					t.Fatalf("Wanted error containing %q but got error %q from spec: %s", test.WantErr, err, specjson)
				}
			} else if test.WantErr != "" {
				t.Fatalf("Wanted error %q but got none from spec: %s", err, specjson)
			}
			if dsn != test.WantDSN {
				t.Fatalf("Wanted DSN:\n\t%q\nbut got\n\t%q\nfrom spec:\n\t%s", test.WantDSN, dsn, specjson)
			}
		})
	}
}

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, client.JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "spec with connection_string",
			Spec: `{"connection_string": "conn"}`,
		},
		{
			Name: "spec with connection_string and pk",
			Spec: `{"connection_string": "conn", "private_key":"foo"}`,
		},
		{
			Name: "spec with bool connection_string",
			Spec: `{"connection_string": true, "private_key":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with null connection_string",
			Spec: `{"connection_string": null, "private_key":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with int connection_string",
			Spec: `{"connection_string": 123, "private_key":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"connection_string": "abc", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"connection_string": "abc", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"connection_string": "abc",  "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"connection_string": "abc", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"connection_string": "abc", "unknown": "test"}`,
			Err:  true,
		},
	})
}
