set -x
set -e

for d in ./resources/services/*/ ; do
  # check if there is a new cq-gen .hcl config file
  if git diff --name-status origin/main HEAD -- $d | grep -q '^A.*\.hcl$'; then
    # .hcl file was newly added
    if grep -q '//check-for-changes' "$d"*.hcl; then
      echo "cq-gen config files in $d are OK";
    else
      echo "//check-for-changes must be present in all new cq-gen config files";
      exit 1;
    fi;
  else
    echo "No new cq-gen config files found in $d"
  fi;
done