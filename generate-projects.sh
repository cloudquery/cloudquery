#!/bin/bash

i=51

until [ $i -gt 100 ]
do
  echo i: $i;
  gcloud projects create test-bernays-$(echo $i) --folder=154109169532 --enable-cloud-apis
  
  ((i=i+1));
done

