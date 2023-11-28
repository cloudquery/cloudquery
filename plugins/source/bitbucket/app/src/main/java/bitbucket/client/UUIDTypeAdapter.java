package bitbucket.client;

import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.util.UUID;

public class UUIDTypeAdapter extends TypeAdapter<UUID> {
  @Override
  public void write(JsonWriter out, UUID value) throws IOException {
    out.value("{" + value.toString() + "}");
  }

  @Override
  public UUID read(JsonReader in) throws IOException {
    String uuidString = in.nextString();
    return UUID.fromString(uuidString.replace("{", "").replace("}", ""));
  }
}
