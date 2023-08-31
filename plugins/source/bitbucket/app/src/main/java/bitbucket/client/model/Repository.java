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
public class Repository {
  @JsonAdapter(UUIDTypeAdapter.class)
  private UUID uuid;

  private String type;

  @SerializedName("full_name")
  private String fullName;

  @SerializedName("is_private")
  private boolean isPrivate;

  private Repository parent;
  private String scm;
  private String name;
  private String description;

  @SerializedName("created_on")
  private LocalDateTime createdOn;

  @SerializedName("updated_on")
  private LocalDateTime updatedOn;

  private int size;
  private String language;

  @SerializedName("has_issues")
  private boolean hasIssues;

  @SerializedName("has_wiki")
  private boolean hasWiki;

  @SerializedName("fork_policy")
  private String forkPolicy;

  private Object project;

  @SerializedName("mainbranch")
  private Object mainBranch;

  private Map<String, Object> links = new HashMap<>();
}
