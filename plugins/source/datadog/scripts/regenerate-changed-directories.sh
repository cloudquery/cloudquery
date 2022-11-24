set -x
set -e

for d in ./resources/services/*/ ; do
  # check whether directory changed in this branch
  if git diff --quiet origin/main HEAD -- $d; then
    echo "no changes in $d";
    continue;
  fi

  # regenerate if //check-for-changes is present in an .hcl file
  if grep -s -q '//check-for-changes' "$d"*.hcl; then
   (cd $d && go generate);
  fi
done