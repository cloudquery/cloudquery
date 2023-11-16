package bitbucket.client.model;

import bitbucket.client.UUIDTypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import lombok.Data;
import lombok.ToString;

@Data
@ToString
public class Workspace {
  @JsonAdapter(UUIDTypeAdapter.class)
  private UUID uuid;

  private String name;
  private String type;
  private String slug;

  @SerializedName("is_private")
  private boolean isPrivate;

  @SerializedName("created_on")
  private LocalDateTime createdOn;

  @SerializedName("updated_on")
  private LocalDateTime updatedOn;

  private Map<String, Object> links = new HashMap<>();
}
