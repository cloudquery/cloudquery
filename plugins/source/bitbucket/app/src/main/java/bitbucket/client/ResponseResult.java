package bitbucket.client;

import java.util.ArrayList;
import java.util.List;
import lombok.Data;
import lombok.ToString;

@Data
@ToString
public class ResponseResult<T> {
  private int page;
  private int size;
  private int pagelen;
  private String next;
  private String previous;
  private List<T> values = new ArrayList<>();
}
