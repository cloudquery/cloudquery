package bitbucket.client.configuration;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

public class SpecTest {
  @Test
  public void shouldThrowExceptionIfWeDoNotHaveRequiredConfiguration() {
    Spec spec = new Spec();

    assertThrows(ConfigurationException.class, spec::validate);

    spec.setUsername("username");
    spec.setPassword("password");

    assertDoesNotThrow(spec::validate);
  }

  @Test
  public void shouldNotThrowExceptionIfRequiredParametersAreSet() {
    Spec spec = new Spec();

    spec.setUsername("username");
    spec.setPassword("password");

    assertDoesNotThrow(spec::validate);
  }
}
