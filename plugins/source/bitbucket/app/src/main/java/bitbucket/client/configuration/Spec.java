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
}
