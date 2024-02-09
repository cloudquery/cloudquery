package bitbucket.client.configuration;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.Data;

@Data
public class Spec {
  public static int DEFAULT_CONCURRENCY = 1000;

  // Required configuration
  private String username;
  private String password;

  // Optional configuration
  private int concurrency = DEFAULT_CONCURRENCY;

  public static Spec fromJSON(String spec) throws JsonProcessingException {
    return new ObjectMapper().readValue(spec, Spec.class);
  }

  public void validate() throws ConfigurationException {
    if (username == null || password == null) {
      throw new ConfigurationException("Username and password are required");
    }
  }

  public static final String jsonSchema =
      """
          {
            "$schema": "https://json-schema.org/draft/2020-12/schema",
            "$id": "https://github.com/cloudquery/cloudquery/plugins/source/bitbicket/spec",
            "$ref": "#/$defs/Spec",
            "$defs": {
              "Spec": {
                "properties": {
                  "username": {
                    "type": "string",
                    "minLength": 1,
                    "description": "BitBucket user."
                  },
                  "password": {
                    "type": "string",
                    "minLength": 1,
                    "description": "BitBucket password."
                  },
                },
                "additionalProperties": false,
                "type": "object",
                "required": [
                  "username",
                  "password"
                ]
              },
            }
          }
          """;
}
